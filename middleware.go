package emas

import (
	"encoding/json"
	"bytes"
	"io"
	"strings"

	"emas/lib/request"
	"emas/lib/response"
	"emas/lib/endpoint"
)

type Middleware struct {
	Client Client
}

func (c *Middleware) Call(method, path string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = "http://php7.orori.com:11029/v1" + path
	if c.Client.Environment == "prod" {
		path = "https://owa.e-mas.com/v1" + path
	}

	return c.Client.Call(method, path, body)
}

func (g *Middleware) TransactionInit(req *request.TransactionInit) (response.Transaction, response.Error, error) {
	trans 		:= response.Transaction{}
	error 		:= response.Error{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", endpoint.BuyInit, bytes.NewBuffer(jsonReq))
	json.Unmarshal(body, &trans)
	json.Unmarshal(body, &error)
	if err != nil {
		g.Client.Logger.Println("Error charging: ", err)
		return trans, error, err
	}

	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
	}

	return trans, error, nil
}