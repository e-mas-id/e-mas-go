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
	
	res,err := midleware.ProductList()
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)

}