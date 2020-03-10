package emas

type ReqTransactionInit struct {
	Customer 	             TransCustomer 	`json:"customer"`
	Email                    string         `json:"email"`
	IdCardNumber             string         `json:"id_card_number"`
	IdCardPhotoUrl           string         `json:"id_card_photo_url"`
	IdCardPhotoSelfieUrl     string         `json:"id_card_photo_selfie_url"`
	Gender                   string         `json:"gender"`
	Weight 		             float64 	    `json:"weight"`
	Amount                   int            `json:"amount"`
	Insurance 	             int 		    `json:"insurance"`
	Birthdate                string         `json:"birthdate"`
	TransactionWithdraw
}

type ReqTransactionConfirm struct {
	MerchantCustomerId  string 		`json:"merchant_customer_id"`
	PaymentId 			string 		`json:"payment_id"`
	VendorTransactionId string 		`json:"vendor_transaction_id"`
}

type ReqTransactionCancel struct {
	MerchantCustomerId  string 		`json:"merchant_customer_id"`
	PaymentId 			string 		`json:"payment_id"`
	VendorTransactionId string 		`json:"vendor_transaction_id"`
}

type TransCustomer struct {
	MerchantCustomerId       string     `json:"merchant_customer_id"`
	Handphone		         string		`json:"handphone"`
	Fullname 		         string		`json:"fullname"`
	Email                    string     `json:"email"`
	IdCardNumber             string     `json:"id_card_number"`
	IdCardPhotoUrl           string     `json:"id_card_photo_url"`
	IdCardPhotoSelfieUrl     string     `json:"id_card_photo_selfie_url"`
	Gender                   string     `json:"gender"`
	Birthdate                string     `json:"birthdate"`
}

type  ProductLogReq struct {
	Page        int     `json:"page"`
	Type        string  `json:"type"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Limit       int     `json:"limit"`
}

type TransactionListReq struct {
	Page                int     `json:"page"`
	Offset              int     `json:"offset"`
	Limit               int     `json:"limit"`
	Status              int     `json:"status"`
	StartDate           string  `json:"start_date"`
	EndDate             string  `json:"end_date"`
	OrderBy             string  `json:"order_by"`
	MerchantCustomerId  string  `json:"merchant_customer_id"`
}