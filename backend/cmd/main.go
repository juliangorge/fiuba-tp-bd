package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"bdd-back/employees"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	employeesRepo := setupPostgresConnection()

	collection, client, mongoDBContext, cancel := setupMongoDBConnection()
	defer closeMongoDBConnection(client, mongoDBContext, cancel)

	// This chunk is only for the initial MongoDB setup
	var result bson.M
	err := collection.FindOne(mongoDBContext, bson.D{}).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}
	fmt.Printf("Found document: %v\n", result)

	router := http.NewServeMux()
	router.HandleFunc("/ping", ping)

	employeeController := employees.NewEmployeeController(employeesRepo)

	router.HandleFunc("/employees", employeeController.GetAll)
	router.HandleFunc("/employees/{id}", employeeController.GetByID)
	router.HandleFunc("POST /employees", employeeController.Create)
	router.HandleFunc("PUT /employees/{id}", employeeController.Update)
	router.HandleFunc("DELETE /employees/{id}", employeeController.Delete)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Server is running on port 8080")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}

func setupPostgresConnection() *employees.EmployeeSQLStorage {
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	employeesRepo := employees.NewEmployeeSQLStorage(db)
	return employeesRepo
}

func setupMongoDBConnection() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	mongoDBHost := os.Getenv("MONGODB_HOST")
	mongoDBUser := os.Getenv("MONGODB_USER")
	mongoDBPassword := os.Getenv("MONGODB_PASSWORD")
	mongoDBDatabase := "local"
	mongoDBCollection := "startup_log"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongoDBUser+":"+mongoDBPassword+"@"+mongoDBHost+":27017"))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}
	collection := client.Database(mongoDBDatabase).Collection(mongoDBCollection)
	return collection, client, ctx, cancel
}

func closeMongoDBConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}
