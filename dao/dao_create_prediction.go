package dao

import (
    "context"
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"

)

func DB_CreatePrediction(object *dto.ForecastPrediction) error {

	_, err := dbConfig.DATABASE.Collection("Predictions").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}