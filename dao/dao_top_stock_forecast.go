package dao

import (
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func DB_Top5ProductsByStockForForecast() ([]dto.TopProductStockForForecast, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
			"$lookup": bson.M{
				"from":         "Products",
				"localField":   "productId",
				"foreignField": "productId",
				"as":           "product",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$product",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":              0,
				"productId":        1,
				"productName":      1,
				"currentInventory": 1,

				"price": bson.M{
					"$ifNull": bson.A{
						"$product.price",
						0,
					},
				},

				"discount": bson.M{
					"$ifNull": bson.A{
						"$product.discountPercent",
						0,
					},
				},

				"category": bson.M{
					"$ifNull": bson.A{
						"$product.category",
						"",
					},
				},

				"storeId": bson.M{
					"$ifNull": bson.A{
						"$product.storeId",
						"$storeId",
					},
				},
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := []dto.TopProductStockForForecast{}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}