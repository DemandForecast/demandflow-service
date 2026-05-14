package dto

type Inventory struct {
	InventoryId string `json:"InventoryId"`
	ProductId   string `json:"ProductId"`
	StoreID     string `json:"StoreID"`
	CurrentInventory  int `json:"CurrentInventory"`
	UnitsSold         int `json:"UnitsSold"`
	UnitsOrdered      int `json:"UnitsOrdered"`
	MaximumStockLevel int `json:"maximum_stock_level"`
	DemandForecast float64 `json:"DemandForecast"`
	LastRestockedDate string `json:"LastRestockedDate"`
	IsLowStock bool `json:"IsLowStock"`
	Deleted    bool `json:"deleted"`

	Base
}
