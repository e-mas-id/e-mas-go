package main

import (
	"fmt"
	emas "e-mas-go"
)

func main(){
	
	client             := emas.NewClient()
	client.AppId        = "TOKOPEDIA"
	client.SecretKey    = "1b111ff465886f7220972d70cbf010f90d372414"
	client.Environment  = "dev"
	client.Debug        = true
	
	midleware          := emas.Middleware{
		Client:client,
	}
	
	var trans_detail emas.TransactionWithdrawDetail
	var arr_trans_detail []*emas.TransactionWithdrawDetail
	
	trans_detail.Id = 1
	trans_detail.Count = 1
	
	arr_trans_detail = append(arr_trans_detail,&trans_detail)
	
	req := &emas.ReqTransactionInit{
		Customer:
		emas.TransCustomer{
			MerchantCustomerId      :"1234567",
			Fullname                :"Indra Octama",
			Email                   :"indra.octama@orori.com",
			Gender                  :"male",
			Handphone               :"082111833438",
			Birthdate               :"1970-01-01",
			IdCardNumber            :"123456789",
			IdCardPhotoSelfieUrl    :"https://example.com/selfie.jpg",
			IdCardPhotoUrl          :"https://example.com/photo.jpg",
		},
		Weight:1,
		TransactionWithdraw:
		emas.TransactionWithdraw{
			Address:
			&emas.TransactionWithdrawAddress{
				Address:"Jln. Abdul Muis No.46",
				City:"Bandung",
				Province:"Jawa Barat",
				Zipcode:"10160",
			},
			ShippingCompany: 1 , //1 : JNE REG ; 2 : JNE YES ; 12 : RPX
			Insurance     : 1, //1 : Yes ; 0 ; No
			Detail: arr_trans_detail,
		},
	}
	
	res,err := midleware.WithdrawInit(req)
	if err != nil{
		fmt.Println(err.Error())
	}
	
	fmt.Println(res)
}