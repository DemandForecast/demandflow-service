package dao

import (
	"Suppliers/dbConfig"
	"Suppliers/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindCustomerbyCustomerId (CustomerId string) (*dto.Customer, error) {
	var object dto.Customer

	err := dbConfig.DATABASE.Collection("Customers").FindOne(context.Background(), bson.M{"customerid": CustomerId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
