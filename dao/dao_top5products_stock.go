package dao

import (
	"DemandFlow-Service/dbConfig"
	"context"
	"DemandFlow-Service/dto"

	"go.mongodb.org/mongo-driver/bson"
)


func DB_Top5ProductsByStock() ([]dto.TopProductStock, error) {

	collection := dbConfig.DATABASE.Collection("Inventory")

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"deleted": false,
			},
		},
		{
			"$sort": bson.M{
				"currentInventory": -1,
			},
		},
		{
			"$limit": 5,
		},
		{
			"$project": bson.M{
				"_id":              0,
				"productId":        1,
				"productName":      1,
				"currentInventory": 1,
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	var result []dto.TopProductStock

	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}