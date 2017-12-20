package emas

type TransactionInit struct {
	CustomerId 		string 		`json:"customer_id"`
	Weight 			float64 	`json:"weight"`
}

type TransactionConfirm struct {
	CustomerId 			string 		`json:"customer_id"`
	PaymentId 			string 		`json:"payment_id"`
	VendorTransactionId string 		`json:"vendor_transaction_id"`
}