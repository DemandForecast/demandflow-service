package dao

import (
	"Suppliers/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteCustomer (CustomerId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Customers").UpdateOne(context.Background(), bson.M{"customerid": CustomerId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

