package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database


const DATABASE_URL = "mongodb+srv://ishara_dilshan:WT8P4WSm5MymQ5bz@demandforecastcluster-1.d2z4vqo.mongodb.net/?appName=demandForecastCluster-1"

const DATABASE_NAME ="Demand-Flow-Service-DB"