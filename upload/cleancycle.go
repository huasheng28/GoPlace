package main

import (
	"io/ioutil"
	"fmt"
	"huasheng28/upload"
	"os"
	"strconv"
	"regexp"
)

func getImgNames(imgFolder string) []string {
	var imgNames []string
	files, _ := ioutil.ReadDir(imgFolder)
	for _,file := range files{
		imgNames = append(imgNames, file.Name())
	}
	return imgNames
}


func cleanCycle(csvName string) {
	url := "http://192.168.1.186:8080/v1/identify"
	//url := upload.ApiUrlText(csvName)
	_, _, _, paramName, extraParam, headers := upload.ReadCsvFile(csvName)
	//把图片名称转换为图片地址
	currentPath, _ := os.Getwd()
	//fmt.Println(currentPath)
	imgNames := getImgNames(currentPath+"/image")
	//fmt.Println(imgNames)
	var imgPath []string
	for _, k := range imgNames {
		absPath := currentPath + "/image/" + k
		imgPath = append(imgPath, absPath)
	}
	//开始上传并输出
	var respBody string
	for i, j := range imgPath {
		fmt.Printf("正在测试[" + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(imgNames)) + "]...")
		respBody = upload.DoUpload(url, j, paramName, extraParam, headers)
		if respBody == "timeout" {
			fmt.Println("%v 连接超时重试",imgNames[i])
			aText := imgNames[i] + "连接超时重试"
			upload.WriteFile(aText)
			respBody = upload.DoUpload(url, j, paramName, extraParam, headers)
		}else if ok, err :=regexp.MatchString("\"errcode\":50001",respBody);ok{
			if err !=nil {
				panic(err)
			}
			respBody = upload.DoUpload(url, j, paramName, extraParam, headers)
			bText := imgNames[i] + "返回错误重试"
			upload.WriteFile(bText)
		}
		upload.WriteTimeToFile()
		upload.WriteFile(j)
		upload.WriteFile(respBody)
	}
	totalText := "测试完成"
	fmt.Println(totalText)
	upload.WriteTimeToFile()
	upload.WriteFile(totalText)
}

func main() {
	cleanCycle("http@192%168%1%1868080+v1+identify")
}
