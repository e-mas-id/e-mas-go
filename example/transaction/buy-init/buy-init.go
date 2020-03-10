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
	
	req := &emas.ReqTransactionInit{
		Customer:
		emas.TransCustomer{
			MerchantCustomerId      :"{YOUR-CUSTOMER-ID}",
			Fullname                :"Indra Octama",
			Email                   :"indra.octama@orori.com",
			Gender                  :"Male",
			Handphone               :"",
			Birthdate               :"",
			IdCardNumber            :"",
			IdCardPhotoSelfieUrl    :"",
			IdCardPhotoUrl          :"",
		},
		Weight:2,
	}
	
	res,err := midleware.BuyInit(req)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
}