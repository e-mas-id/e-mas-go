package response

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