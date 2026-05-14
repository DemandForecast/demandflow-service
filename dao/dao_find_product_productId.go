package dao

import (
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindProductbyProductId (ProductId string) (*dto.Product, error) {
	var object dto.Product

	err := dbConfig.DATABASE.Collection("Products").FindOne(context.Background(), bson.M{"productid": ProductId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
