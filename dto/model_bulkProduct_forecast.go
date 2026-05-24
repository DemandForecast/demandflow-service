package dto

type TopStockForecastRequest struct {
	FromDate string `json:"fromDate" bson:"fromDate"`
	ToDate   string `json:"toDate" bson:"toDate"`

	Weather     string `json:"weather" bson:"weather"`
	Region      string `json:"region" bson:"region"`
	Seasonality string `json:"seasonality" bson:"seasonality"`
	Holiday     bool   `json:"holiday" bson:"holiday"`
}

type TopProductStockForForecast struct {
	ProductID        string `json:"productId" bson:"productId"`
	ProductName      string `json:"productName" bson:"productName"`
	Image            string `json:"image" bson:"image"`
	CurrentInventory int    `json:"currentInventory" bson:"currentInventory"`
	Price    float64 `json:"price" bson:"price"`
	Discount float64 `json:"discount" bson:"discount"`
	Category string  `json:"category" bson:"category"`
	StoreID  string  `json:"storeId" bson:"storeId"`
}
