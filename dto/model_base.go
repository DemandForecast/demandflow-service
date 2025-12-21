package dto

import "time"

type Base struct {
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	LastUpdatedBy string    `json:"lastUpdatedBy"`
}

type Products struct {
	ProductId   string ` json:"ProductId" `
	Quantity   int ` json:"Quantity" `
}