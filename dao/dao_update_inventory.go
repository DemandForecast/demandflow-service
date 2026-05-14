package dao

import (
	"DemandFlow-Service/dbConfig"
    "DemandFlow-Service/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateInventory (object *dto.Inventory)  error {

	result, err := dbConfig.DATABASE.Collection("Inventory").UpdateOne(context.Background(), bson.M{"inventoryid": object.InventoryId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}