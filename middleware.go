package emas

import (
	"encoding/json"
	"bytes"
	"io"
	"strings"
	"errors"
)

type Middleware struct {
	Client Client
}

func (c *Middleware) Call(method, path string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = "http://php7.orori.com:11029/v1/thirdparty" + path
	if c.Client.Environment == "prod" {
		path = "" + path
	}

	return c.Client.Call(method, path, body)
}

func (g *Middleware) BuyInit(req *ReqTransactionInit) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointBuyInit, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error buy init: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}

func (g *Middleware) BuyConfirm(req *ReqTransactionConfirm) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointBuyConfirm, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error buy confirm: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}

func (g *Middleware) BuyCancel(req *ReqTransactionConfirm) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointBuyCancel, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error buy cancel: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}

func (g *Middleware) SellInit(req *ReqTransactionInit) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointSellInit, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error sell init: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}

func (g *Middleware) SellConfirm(req *ReqTransactionConfirm) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointSellConfirm, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error sell reverse: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}

func (g *Middleware) SellReverse(req *ReqTransactionConfirm) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)

	body, err := g.Call("POST", EndpointSellReversal, bytes.NewBuffer(jsonReq))
	if err != nil {
		g.Client.Logger.Println("Error sell reverse: ", err)
		return resp, err
	}

	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &error)
	if error.ErrorMessage != "" {
		g.Client.Logger.Println(error.ErrorMessage)
		return resp, errors.New(error.ErrorMessage)
	}

	return resp, nil
}