package dao

import (
    "context"
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"

)

func DB_CreateProduct(object *dto.Product) error {

	_, err := dbConfig.DATABASE.Collection("Products").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}