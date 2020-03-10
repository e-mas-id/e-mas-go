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
	
	payment_id := "aa301507ebcfe367006a086abe99d211abae0635"
	
	res,err := midleware.TransactionDetail(payment_id)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
	
}
