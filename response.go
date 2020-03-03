package emas

import "time"

type SuccessResponse struct {
	Code 			int			`json:"code"`
	Type			string		`json:"type"`
	Data			interface{}	`json:"data"`
	Total			int			`json:"total"`
	FirstPage		string		`json:"first_page"`
	LastPage		string		`json:"last_page"`
	NextPage		string		`json:"next_page"`
	Timestamp 		string		`json:"timestamp"`
}

type ErrorResponse struct {
	ErrorCode 		int			`json:"error_code"`
	ErrorMessage	string		`json:"error_message"`
	ErrorData		interface{}	`json:"error_data"`
	Timestamp 		string		`json:"timestamp"`
}

type Customer struct {
	Balance 		float64		`json:"balance"`
	Address			string		`json:"address"`
	City			string		`json:"city"`
	Province		string		`json:"province"`
	Zipcode 		string		`json:"zipcode"`
}

type Product struct {
	Id       		uint64		`json:"id"`
	Name 			string		`json:"name"`
	Description		string		`json:"description"`
	Weight			int			`json:"weight"`
	SellPrice		uint32		`json:"sell_price"` // We sell to customer
	BuyPrice		uint32		`json:"buy_price"`  // We buy from customer
	AdminFee		int			`json:"admin_fee"`
	Image			string		`json:"image"`
}

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
	Detail			[]*TransactionWithdrawDetail	`json:"detail"`
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