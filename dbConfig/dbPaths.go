package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database


const DATABASE_URL = "mongodb+srv://pmchamoth_db_user:t0rgjOG2h89qexQD@imomgt.y8j6wpi.mongodb.net/?appName=IMOMgt"

const DATABASE_NAME ="Demand-Forecast"