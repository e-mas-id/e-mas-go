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
	
	req := &emas.ProductLogReq{
		Page:1,
		StartDate:"2020-03-03 00:00:00", //format must be YYYY-MM-DD HH:ii:ss
		EndDate:"2020-03-06 23:59:59",   //format must be YYYY-MM-DD HH:ii:ss
		Limit:30,
		Type:"day",
	}
	
	res,err := midleware.ProductLog(req)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
	
}