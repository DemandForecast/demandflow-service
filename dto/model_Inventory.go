package dto

type Inventory struct {
	InventoryId string `json:"inventoryId" bson:"inventoryId"`
	ProductId   string `json:"productId" bson:"productId"`
	ProductName string `json:"productName" bson:"productName"`
	StoreID     string `json:"storeId" bson:"storeId"`

	CurrentInventory  int     `json:"currentInventory" bson:"currentInventory"`
	UnitsSold         int     `json:"unitsSold" bson:"unitsSold"`
	UnitsOrdered      int     `json:"unitsOrdered" bson:"unitsOrdered"`
	MaximumStockLevel int     `json:"maximumStockLevel" bson:"maximumStockLevel"`
	DemandForecast    float64 `json:"demandForecast" bson:"demandForecast"`

	LastRestockedDate string `json:"lastRestockedDate" bson:"lastRestockedDate"`

	IsLowStock bool `json:"isLowStock" bson:"isLowStock"`
	Deleted    bool `json:"deleted" bson:"deleted"`

	Base
}
