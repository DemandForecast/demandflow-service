package dto

type Customer struct {

CustomerId	string ` json:"CustomerId" `
	ResellerId    string ` json:"ResellerId" `
	MerchantId    string ` json:"MerchantId" `
	IMOId		 string  ` json:"IMOId" `
	Name          string ` json:"Name" `
	ProImg		string ` json:"ProImg" `
	ContactNumber string ` json:"ContactNumber" `
	Email         string ` json:"Email" `
	Address       string ` json:"Address" `
	Status        string ` json:"Status" `
	NIC          string ` json:"NIC" `
	Date string ` json:"Date" `
	DiamondStock int64 ` json:"DiamondStock" `
	Deleted       bool   `json:"deleted"`
	Base
}
