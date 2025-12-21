package dao

import (
	"Suppliers/dbConfig"
    "Suppliers/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateCustomer (object *dto.Customer)  error {

	result, err := dbConfig.DATABASE.Collection("Customers").UpdateOne(context.Background(), bson.M{"customerid": object.CustomerId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}