package dao

import (
	"DemandFlow-Service/dbConfig"
    "DemandFlow-Service/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateProduct (object *dto.Product)  error {

	result, err := dbConfig.DATABASE.Collection("Products").UpdateOne(context.Background(), bson.M{"productid": object.ProductId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}