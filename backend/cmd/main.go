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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	employeetags "bdd-back/employee_tags"
	"bdd-back/employees"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	employeesRepo := setupPostgresConnection()

	runMongoDBMigrations()
	employeeTagStorage, client, mongoDBContext, cancel := setupMongoDBConnection()
	defer closeMongoDBConnection(client, mongoDBContext, cancel)

	router := http.NewServeMux()
	router.HandleFunc("/ping", ping)

	employeeController := employees.NewEmployeeController(employeesRepo)
	employeeTagController := employeetags.NewEmployeeTagController(employeeTagStorage)

	router.HandleFunc("/employees", employeeController.GetAll)
	router.HandleFunc("/employees/{id}", employeeController.GetByID)
	router.HandleFunc("POST /employees", employeeController.Create)
	router.HandleFunc("PUT /employees/{id}", employeeController.Update)
	router.HandleFunc("DELETE /employees/{id}", employeeController.Delete)

	router.HandleFunc("/employee_tags/{employee_id}", employeeTagController.GetAllTagsByID)
	router.HandleFunc("POST /employee_tags/{employee_id}/tags/{tag_name}", employeeTagController.InsertTag)
	router.HandleFunc("DELETE /employee_tags/{employee_id}/tags/{tag_name}", employeeTagController.RemoveTag)

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

func setupMongoDBConnection() (employeetags.EmployeeTagStorage, *mongo.Client, context.Context, context.CancelFunc) {
	mongoDBDatabase := "admin"
	mongoDBCollection := "employee_tags"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(buildMongoDBUriWithAuth()))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	collection := client.Database(mongoDBDatabase).Collection(mongoDBCollection)
	employeeTagsRepo := employeetags.NewEmployeeTagMongoDBStorage(collection)

	return employeeTagsRepo, client, ctx, cancel
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

// Valid operations can be gotten from https://github.com/golang-migrate/migrate/tree/master/database/mongodb/examples/migrations
func runMongoDBMigrations() {
	// Replace with your MongoDB connection string and migrations folder path
	m, err := migrate.New(
		"file://mongodbmigrations",
		buildMongoDBUriWithAuth()+"/admin",
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func buildMongoDBUriWithAuth() string {
	mongoDBHost := os.Getenv("MONGODB_HOST")
	mongoDBUser := os.Getenv("MONGODB_USER")
	mongoDBPassword := os.Getenv("MONGODB_PASSWORD")

	return "mongodb://" + mongoDBUser + ":" + mongoDBPassword + "@" + mongoDBHost + ":27017"
}
