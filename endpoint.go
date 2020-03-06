package emas

const (
	// GET
	EndpointCustomerDetail 		= "/customer/"
	EndpointProductList 		= "/product"
	EndpointProductLog   		= "/product-log"
	EndpointTransaction 		= "/transaction"
	EndpointTransactionDetail 	= "/transaction/"

	// POST
	EndpointBuyInit 			= "/buy-gold"
	EndpointBuyConfirm			= "/buy-gold-confirm"
	EndpointBuyCancel			= "/buy-gold-cancel"
	EndpointSellInit			= "/sell-gold"
	EndpointSellConfirm			= "/sell-gold-confirm"
	EndpointSellReversal		= "/sell-gold-reverse"
	EndpointWithdrawInit		= "/withdraw-gold"
	EndpointWithdrawConfirm		= "/withdraw-gold-confirm"
	EndpointWithdrawCancel		= "/withdraw-cancel"
)