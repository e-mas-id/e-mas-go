package emas

type ReqTransactionInit struct {
	CustomerId 		string 		`json:"customer_id"`
	Weight 			float64 	`json:"weight"`
}

type ReqTransactionConfirm struct {
	CustomerId 			string 		`json:"customer_id"`
	PaymentId 			string 		`json:"payment_id"`
	VendorTransactionId string 		`json:"vendor_transaction_id"`
}