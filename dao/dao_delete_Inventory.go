package dao

import (
	"DemandFlow-Service/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteInventory (InventoryId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Inventory").UpdateOne(context.Background(), bson.M{"inventoryid": InventoryId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

