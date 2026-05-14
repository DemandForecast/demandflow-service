package dto

type Product struct {
	ProductId       string  `json:"ProductId"`
	ProductName     string  `json:"ProductName"`
	Category        string  `json:"Category"`
	Brand           string  `json:"Brand"`
	SKU             string  `json:"sku"`
	Description     string  `json:"Description"`
	Price           float64 `json:"Price"`
	DiscountPercent float64 `json:"discount_percent"`

	IsPerishable bool   `json:"is_perishable"`
	StoreID      string `json:"store_id"`
	SupplierID   string `json:"supplier_id"`
	// Quantity     int    ` json:"Quantity" `
	Deleted       bool   `json:"deleted"`
	Base
}
