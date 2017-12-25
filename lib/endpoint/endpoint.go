package endpoint

const (
	// GET
	CustomerDetail 		= "/customer/"
	Product 			= "/product"
	Transaction 		= "/transaction"
	TransactionDetail 	= "/transaction/"

	// POST
	BuyInit 			= "/buy-gold"
	BuyConfirm			= "/buy-gold-confirm"
	BuyCancel			= "/buy-gold-cancel"
	SellInit			= "/sell-gold"
	SellConfirm			= "/sell-gold-confirm"
	SellReversal		= "/sell-gold-reverse"
	WithdrawInit		= "/withdraw-gold"
	WithdrawConfirm		= "/withdraw-gold-confirm"
)