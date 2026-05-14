package dao

import (
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindInventorybyInventoryId (InventoryId string) (*dto.Inventory, error) {
	var object dto.Inventory

	err := dbConfig.DATABASE.Collection("Inventory").FindOne(context.Background(), bson.M{"inventoryid": InventoryId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
