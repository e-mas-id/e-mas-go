package response

import "time"

type Transaction struct {
	PaymentId 		string 					`json:"payment_id"`
	Description 	string 					`json:"description"`
	Balance 		float64					`json:"balance,omitempty"`
	Weight 			float64 				`json:"weight"`
	Price 			float64 				`json:"price"`
	AdminFee		int						`json:"admin_fee"`
	PricePerGram 	float64 				`json:"price_per_gram"`
	Status 			uint8 					`json:"status"`
	PaymentStatus 	int 					`json:"payment_status"`
	Withdraw		*TransactionWithdraw	`json:"withdraw_detail,omitempty"`
	CreatedAt 		time.Time 				`json:"created_at"`
}

type TransactionWithdraw struct {
	Handphone		string							`json:"handphone"`
	Fullname 		string							`json:"fullname"`
	Address  		*TransactionWithdrawAddress		`json:"address"`
	ShippingCompany int 							`json:"shipping_company"` // 1 = JNE_REG, 2 = JNE_YES, 12 = RPX_NDP
	ShippingFee		int								`json:"shipping_fee"`
	ShippingCode	string							`json:"shipping_code"`
	Detail			[]*TransactionWithdrawDetail		`json:"detail"`
}

type TransactionWithdrawAddress struct {
	Address  		string	`orm:"size(255)" json:"address"`
	City	 		string	`orm:"size(100)" json:"city"`
	Province		string	`orm:"size(100)" json:"province"`
	Zipcode			string	`orm:"size(10)" json:"zipcode"`
}

type TransactionWithdrawDetail struct {
	Id 		uint64 			`json:"id"`
	Count	int				`json:"count"`
}