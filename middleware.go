package emas

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"
)

type Middleware struct {
	Client Client
}

const EMAS_DEVELOPMENT   = "https://oroconnect-dev.e-mas.com/v2/thirdparty"
const EMAS_PRODUCTION    = ""

func (c *Middleware) Call(method, path string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}


	path = EMAS_DEVELOPMENT + path
	if c.Client.Environment == "prod" {
		path = EMAS_PRODUCTION + path
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

func (g *Middleware) BuyCancel(req *ReqTransactionCancel) (SuccessResponse, error) {
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

func (g *Middleware) SellReverse(req *ReqTransactionCancel) (SuccessResponse, error) {
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


func (g *Middleware) WithdrawInit(req *ReqTransactionInit) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)
	
	body, err := g.Call("POST", EndpointWithdrawInit, bytes.NewBuffer(jsonReq))
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


func (g *Middleware) WithdrawConfirm(req *ReqTransactionConfirm) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)
	
	body, err := g.Call("POST", EndpointWithdrawConfirm, bytes.NewBuffer(jsonReq))
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

func (g *Middleware) WithdrawCancel(req *ReqTransactionCancel) (SuccessResponse, error) {
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	jsonReq, _ 	:= json.Marshal(req)
	
	body, err := g.Call("POST", EndpointWithdrawCancel, bytes.NewBuffer(jsonReq))
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

func (g *Middleware) ProductList()(SuccessResponse,error){
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	
	body, err := g.Call("GET", EndpointProductList, bytes.NewBuffer([]byte("")))
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

func (g *Middleware) ProductLog(req *ProductLogReq)(SuccessResponse,error){
	resp 		:= SuccessResponse{}
	error 		:= ErrorResponse{}
	
	extraParam := "?param=1"
	
	if req.Type != ""{
		extraParam += "&type="+req.Type
	}
	
	if req.StartDate != ""{
		extraParam += "&start_date="+url.QueryEscape(req.StartDate)
	}
	
	if req.EndDate != ""{
		extraParam += "&end_date="+url.QueryEscape(req.EndDate)
	}
	
	if req.Page != 0 {
		extraParam += "&page="+strconv.Itoa(req.Page)
	}
	
	if req.Limit != 0 {
		extraParam += "&limit="+strconv.Itoa(req.Limit)
	}
	
	body, err := g.Call("GET", EndpointProductLog+extraParam, bytes.NewBuffer([]byte("")))
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

func (g *Middleware) CustomerProfile(merchant_customer_id string)(resp SuccessResponse,err error){
	
	body, err := g.Call("GET", EndpointCustomerDetail+"/"+merchant_customer_id, bytes.NewBuffer([]byte("")))
	if err != nil {
		g.Client.Logger.Println("Error sell init: ", err)
		return resp, err
	}
	
	json.Unmarshal(body, &resp)
	json.Unmarshal(body, &err)
	if err != nil {
		g.Client.Logger.Println(err.Error())
		return resp, errors.New(err.Error())
	}
	
	return resp, nil
}