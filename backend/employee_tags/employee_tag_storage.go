package employee_tags

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeTagMongoDBStorage struct {
	EmplyeeTagCollection *mongo.Collection
}

func NewEmployeeTagMongoDBStorage(emplyeeTagCollection *mongo.Collection) *EmployeeTagMongoDBStorage {
	return &EmployeeTagMongoDBStorage{EmplyeeTagCollection: emplyeeTagCollection}
}

func (s *EmployeeTagMongoDBStorage) GetAllTagsByID(ctx context.Context, employeeID int) (EmployeeTag, error) {
	query := bson.M{
		"employee_id": employeeID,
	}

	cursor, err := s.EmplyeeTagCollection.Find(ctx, query)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return EmployeeTag{}, err
	}
	defer cursor.Close(ctx)

	var results []EmployeeTag
	if err = cursor.All(ctx, &results); err != nil {
		log.Printf("Failed to decode query results: %v\n", err)
		return EmployeeTag{}, err
	}

	if len(results) == 0 {
		return EmployeeTag{
			EmployeeID: employeeID,
			Tags:       []string{},
		}, nil
	}

	return results[0], nil
}

func (s *EmployeeTagMongoDBStorage) InsertTag(ctx context.Context, employeeID int, tagToAdd string) error {
	update := bson.M{
		"$addToSet": bson.M{
			"tags": tagToAdd,
		},
	}

	filter := bson.M{
		"employee_id": employeeID,
	}

	opts := options.Update().SetUpsert(true)

	result, err := s.EmplyeeTagCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Failed to update document: %v\n", err)
		return err
	}

	log.Printf("Matched %d document(s) and modified %d document(s)\n", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (s *EmployeeTagMongoDBStorage) RemoveTag(ctx context.Context, employeeID int, tagToRemove string) error {
	update := bson.M{
		"$pull": bson.M{
			"tags": tagToRemove,
		},
	}

	filter := bson.M{
		"employee_id": employeeID,
	}

	result, err := s.EmplyeeTagCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to update document: %v\n", err)
		return err
	}

	fmt.Printf("Matched %d document(s) and modified %d document(s)\n", result.MatchedCount, result.ModifiedCount)
	return nil
}
