# e-mas-go
library for thirdparty connect to e-mas API

## Example Usage

```go
    cli := emas.NewClient()
    cli.AppId 		= "APPLICATION ID FROM US"
    cli.SecretKey 	= "SECRET KEY FROM US"
    cli.Debug 		= true
    cli.Environment = "dev"

    mid := emas.Middleware{
        Client: cli,
    }

    req := &request.TransactionInit{
        CustomerId	: "1",
        Weight		: 1,
    }

    resp, err := mid.TransactionInit(req)
    if err != nil {
        // do something for error
    }

    // do something with resp
```

You can handle the resp object or extract it into response.Transaction if you want

For more information and support please contact info@e-mas.com or dev@e-mas.com.