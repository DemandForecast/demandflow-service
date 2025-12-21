package dao

import (
	"Suppliers/dbConfig"
	"Suppliers/dto"
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindallCustomer () (*[]dto.Customer, error) {
	var objects []dto.Customer
	results, err := dbConfig.DATABASE.Collection("Customers").Find(context.Background(), bson.M{"deleted":false})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Customer
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Customers")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}


func DB_FindCustomersWithPg( page, size, search string) (int64, *[]dto.Customer, error) {
    pageInt, err := strconv.Atoi(page)
    if err != nil || pageInt < 1 {
        return 0, nil, errors.New("invalid page number")
    }
    sizeInt, err := strconv.Atoi(size)
    if err != nil || sizeInt < 1 {
        return 0, nil, errors.New("invalid page size")
    }
    skip := int64((pageInt - 1) * sizeInt)
    limit := int64(sizeInt)

    filter := bson.M{
        "deleted":        false,
    }
    if search != "" {
        filter["$or"] = []bson.M{
            {"name": bson.M{"$regex": search, "$options": "i"}},
            {"email": bson.M{"$regex": search, "$options": "i"}},
            {"address": bson.M{"$regex": search, "$options": "i"}},
			{"suppliers": bson.M{"$regex": search, "$options": "i"}},
            {"contactperson": bson.M{"$regex": search, "$options": "i"}},
            {"company": bson.M{"$regex": search, "$options": "i"}},
			{"category": bson.M{"$regex": search, "$options": "i"}},

        }
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    pipeline := []bson.M{
        {"$match": filter},
        {"$facet": bson.M{
            "metadata": []bson.M{{"$count": "total"}},
            "data": []bson.M{
                {"$skip": skip},
                {"$limit": limit},
            },
        }},
    }

    cursor, err := dbConfig.DATABASE.Collection("Customers").Aggregate(ctx, pipeline)
    if err != nil {
        return 0, nil, err
    }
    defer cursor.Close(ctx)

    var results struct {
        Metadata []struct {
            Total int32 `bson:"total"`
        } `bson:"metadata"`
        Data []dto.Customer `bson:"data"`
    }

    if cursor.Next(ctx) {
        if err := cursor.Decode(&results); err != nil {
            return 0, nil, err
        }
    }

    var customers []dto.Customer
    if len(results.Data) > 0 {
        customers = results.Data
    } else {
        customers = []dto.Customer{}
    }

    var count int64 = 0
    if len(results.Metadata) > 0 {
        count = int64(results.Metadata[0].Total)
    }

    return count, &customers, nil
}