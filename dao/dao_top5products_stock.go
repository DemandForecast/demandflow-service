package dao

import (
	"DemandFlow-Service/dbConfig"
	"DemandFlow-Service/dto"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func DB_Top5ProductsByUnitsSold() ([]dto.Product, error) {
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
				"unitsSold": -1,
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
				"preserveNullAndEmptyArrays": false,
			},
		},
		{
			"$match": bson.M{
				"product.deleted": false,
			},
		},
		{
			"$replaceRoot": bson.M{
				"newRoot": "$product",
			},
		},
		{
			"$project": bson.M{
				"_id":             0,
				"productId":       1,
				"productName":     1,
				"category":        1,
				"brand":           1,
				"sku":             1,
				"description":     1,
				"image":           1,
				"price":           1,
				"discountPercent": 1,
				"isPerishable":    1,
				"storeId":         1,
				"supplierId":      1,
				"quantity":        1,
				"deleted":         1,
				"createdAt":       1,
				"updatedAt":       1,
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []dto.Product

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}