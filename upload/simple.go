package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main() {
	body:=strings.NewReader("{\"instances\": [1.0,2.0,5.0]}")
	req, err := http.NewRequest("POST", "http://192.168.1.134:8501/v1/models/half_plus_three:predict", body)
	if err!=nil{
		fmt.Println(err)
	}
	client:=http.Client{}
	resp, err := client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(respBody)
	//resp, err := http.Post("http://192.168.1.134:8501/v1/models/half_plus_three:predict", "application/x-www-form-urlencoded", body)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//respBody, err := ioutil.ReadAll(resp.Body)
	//
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(respBody)
	//fmt.Println(resp.Body)
}