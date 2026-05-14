package dao

import (
	"DemandFlow-Service/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteProduct (CustomerId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Products").UpdateOne(context.Background(), bson.M{"productid": CustomerId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

