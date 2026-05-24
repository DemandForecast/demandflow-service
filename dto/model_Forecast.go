package dto

type ProductForecastRequest struct {
	ProductID string `json:"productId" bson:"productId"`

	FromDate string `json:"fromDate" bson:"fromDate"`
	ToDate   string `json:"toDate" bson:"toDate"`

	Price     float64 `json:"price" bson:"price"`
	Discount  float64 `json:"discount" bson:"discount"`
	Inventory int     `json:"inventory" bson:"inventory"`

	Weather     string `json:"weather" bson:"weather"`
	Region      string `json:"region" bson:"region"`
	Category    string `json:"category" bson:"category"`
	Seasonality string `json:"seasonality" bson:"seasonality"`
	Holiday      bool   `json:"holiday" bson:"holiday"`
}

type ForecastSummary struct {
	ForecastDays         int     `json:"forecastDays" bson:"forecastDays"`
	TotalPredictedDemand float64 `json:"totalPredictedDemand" bson:"totalPredictedDemand"`
	AverageDailyDemand   float64 `json:"averageDailyDemand" bson:"averageDailyDemand"`
	MinimumDailyDemand   float64 `json:"minimumDailyDemand" bson:"minimumDailyDemand"`
	MaximumDailyDemand   float64 `json:"maximumDailyDemand" bson:"maximumDailyDemand"`
}

type ProductForecastData struct {
	ProductID string `json:"productId" bson:"productId"`
	Image      string `json:"image" bson:"image"`
	FromDate       string  `json:"fromDate" bson:"fromDate"`
	ToDate         string  `json:"toDate" bson:"toDate"`
	PredictedDate  string  `json:"predictedDate" bson:"predictedDate"`
	DemandForecast float64 `json:"demandForecast" bson:"demandForecast"`

	ModelName string `json:"modelName" bson:"modelName"`

	Summary ForecastSummary `json:"summary" bson:"summary"`
}

// type ProductForecastResponse struct {
// 	ForecastID string `json:"forecastId" bson:"forecastId"`

// 	ProductForecastData
// }

type ForecastPrediction struct {
	// ForecastID   string `json:"forecastId" bson:"forecastId"`
	PredictionID string `json:"predictionId" bson:"predictionId"`

	ForecastType string `json:"forecastType" bson:"forecastType"`

	ProductForecastData

	Deleted bool `json:"deleted" bson:"deleted"`

	Base
}
