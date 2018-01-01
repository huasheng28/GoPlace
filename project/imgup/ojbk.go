package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"
)

func check(err error){
	if err!=nil{
		fmt.Println(err)
	}
}

func ReadCsvFile(csvName string) ([]string, []string, int,string,map[string]string) {
	//读取文件
	currentPath, _ := os.Getwd()
	csvFile := currentPath + "/csv/" + csvName + ".csv"
	f, err := os.Open(csvFile)
	check(err)
	defer f.Close()
	r := csv.NewReader(f)
	record, err := r.ReadAll()
	check(err)
	//图片数量
	picNum := len(record)-1
	//上传参数
	paramName:=record[0][0]
	//其他参数
	var extraParam map[string]string
	extraParamJson:=record[0][1]
	if extraParamJson==""{
		extraParam=map[string]string{}
	}else {
		err1 := json.Unmarshal([]byte(extraParamJson), &extraParam)
		check(err1)
	}

	//图片地址及正确返回参数
	var imgName, rightResult []string
	for _, a := range record[1:] {
		imgName = append(imgName, a[0])
		rightResult = append(rightResult, a[1])
	}
	return imgName, rightResult, picNum,paramName,extraParam
}
func main() {
	imgName, rightResult, picNum,paramName,extraParam:=ReadCsvFile("m%seeunsee%cn+intelligent-packaging-check+check%php")
	fmt.Println(imgName)
	fmt.Println(rightResult)
	fmt.Println(picNum)
	fmt.Println(paramName)
	fmt.Println(extraParam)
}
