package dao

import (
	"DemandFlow-Service/dbConfig"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func DB_TotalInventoryStock() (int64, error) {

	collection := dbConfig.DATABASE.Collection("Inventory")

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"deleted": false,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"totalStock": bson.M{
					"$sum": "$currentInventory",
				},
			},
		},
	}

	cursor, err := collection.Aggregate(
		context.Background(),
		pipeline,
	)

	if err != nil {
		return 0, err
	}

	var results []bson.M

	if err = cursor.All(context.Background(), &results); err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, nil
	}

	totalStock := results[0]["totalStock"].(int32)

	return int64(totalStock), nil
}