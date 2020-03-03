# e-mas-go
library for thirdparty connect to e-mas API

## Example Usage
### Buy Init
```go
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
    		//do something when there is an error
    	}
    
    // do something with res
```

You can handle the res object or extract it into emas.Transaction if you want

For more information and support please contact info@e-mas.com or dev@e-mas.com.