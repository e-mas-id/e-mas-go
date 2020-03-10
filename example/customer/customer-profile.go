package main

import (
	emas "e-mas-go"
	"fmt"
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
	
	merchant_customer_id := "1234567"
	
	res,err := midleware.CustomerProfile(merchant_customer_id)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
	
}