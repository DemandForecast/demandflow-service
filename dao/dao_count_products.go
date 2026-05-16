package dao

import (
	"DemandFlow-Service/dbConfig"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func DB_CountProducts() (int64, error) {

	count, err := dbConfig.DATABASE.Collection("Products").CountDocuments(context.Background(), bson.M{"deleted": false},)
	if err != nil {
		return 0, err
	}

	return count, nil
}