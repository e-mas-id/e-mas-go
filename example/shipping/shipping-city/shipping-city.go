package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	
	resp, err := http.Get("https://s3-ap-southeast-1.amazonaws.com/assets.orori.com/indonesia/js/city-new.json")
	if err != nil{
		fmt.Println("Error: "+err.Error())
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Error: "+err.Error())
	}
	
	log.Println(string(body))
}


