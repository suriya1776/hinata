package models

type BankState struct {
	BankName string `bson:"bankName" json:"bankName"`
	// Add other bank-related fields as needed
}
