package dto

type Product struct {
	ProductId       string  `json:"productId" bson:"productId"`
	ProductName     string  `json:"productName" bson:"productName"`
	Category        string  `json:"category" bson:"category"`
	Brand           string  `json:"brand" bson:"brand"`
	SKU             string  `json:"sku" bson:"sku"`
	Description     string  `json:"description" bson:"description"`
	Image           string  `json:"image" bson:"image"`
	Price           float64 `json:"price" bson:"price"`
	DiscountPercent float64 `json:"discountPercent" bson:"discountPercent"`

	IsPerishable bool   `json:"isPerishable" bson:"isPerishable"`
	StoreID      string `json:"storeId" bson:"storeId"`
	SupplierID   string `json:"supplierId" bson:"supplierId"`
	Quantity     int    `json:"quantity" bson:"quantity" `
	Deleted      bool   `json:"deleted" bson:"deleted"`

	Base
}

type TopProductStock struct {
	ProductId        string `bson:"productId" json:"productId"`
	ProductName      string `bson:"productName" json:"productName"`
	CurrentInventory int    `bson:"currentInventory" json:"currentInventory"`
}
