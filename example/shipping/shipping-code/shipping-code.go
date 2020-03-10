package main

import (
	emas "e-mas-go"
	"fmt"
)

func main(){
	client             := emas.NewClient()
	client.Environment  = "dev"
	client.Debug        = true
	
	midleware          := emas.Middleware{
		Client:client,
	}
	
	res,err := midleware.ShippingCode()
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
	
}