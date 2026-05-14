package dto

type Product struct {
	ProductId       string  `json:"ProductId"`
	ProductName     string  `json:"ProductName"`
	Category        string  `json:"Category"`
	Brand           string  `json:"Brand"`
	SKU             string  `json:"sku"`
	Description     string  `json:"Description"`
	Price           float64 `json:"Price"`
	DiscountPercent float64 `json:"DiscountPercent"`
	Image        string  `json:"Image"`
	IsPerishable bool   `json:"IsPerishable"`
	StoreID      string `json:"StoreID"`
	SupplierID   string `json:"SupplierID"`
	Quantity     int    ` json:"Quantity" `
	MOQ		  int    `json:"MOQ"`
	Deleted       bool   `json:"deleted"`
	Base
}
