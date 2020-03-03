package main

import (
	"fmt"
	emas "e-mas-go"
)

func main(){
	
	client             := emas.NewClient()
	client.AppId        = "{YOUR-APP-ID}"
	client.SecretKey    = "{YOUR-SECRET-KEY}"
	client.Environment  = "dev"
	client.Debug        = true
	
	midleware          := emas.Middleware{
		Client:client,
	}
	
	req := &emas.ReqTransactionCancel{
		MerchantCustomerId  :"{YOUR-CUSTOMER-ID}",
		PaymentId           :"{PAYMENT-ID-FROM-BUY-INIT}",
		VendorTransactionId :"{YOUR-UNIQUE-TRANSACTION-ID}",
	}
	
	res,err := midleware.SellReverse(req)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
}