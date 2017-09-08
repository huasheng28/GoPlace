package main

import (

	"fmt"
	"strconv"
	"os"
	"github.com/go-ini/ini"
)

func inputFile(num int) (url string, result string) {
	sec := strconv.Itoa(num)
	inputFilePath, _ := os.Getwd()
	inputFilePath += "/input.ini"
	cfg, err := ini.Load(inputFilePath)
	if err != nil {
		fmt.Println(err)
	}
	url= cfg.Section(sec).Key("url").String()

	result = cfg.Section(sec).Key("result").String()

	return url, result
}
func main()  {

	url,result:=inputFile(1)
	fmt.Println(url,result)
}
