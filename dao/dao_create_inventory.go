package dao

import (
    "context"
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"

)

func DB_CreateInventory(object *dto.Inventory) error {

	_, err := dbConfig.DATABASE.Collection("Inventory").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}