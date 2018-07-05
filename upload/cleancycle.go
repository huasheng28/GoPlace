package main

import (
	"io/ioutil"
	"fmt"
	"upload"
	"os"
	"strconv"
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
	upload.WriteTimeToFile()
	//url := upload.ApiUrlText(csvName)
	//_, _, _, paramName, extraParam, headers := upload.ReadCsvFile(csvName)
	//把图片名称转换为图片地址
	currentPath, _ := os.Getwd()
	fmt.Println(currentPath)
	imgNames := getImgNames(currentPath+"/image")
	fmt.Println(imgNames)
	var imgPath []string
	for _, k := range imgNames {
		absPath := currentPath + "/image/" + csvName + "/" + k
		imgPath = append(imgPath, absPath)
		fmt.Println(imgPath)
	}
	//开始上传并输出
	var respBody string
	var total, timeout = 0, 0
	for i, j := range imgPath {
		fmt.Printf("正在测试[" + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(imgNames)) + "]...")
		respBody = upload.DoUpload(url, j, paramName, extraParam, headers)
		if respBody == "" {
			timeout++
			fmt.Println("连接超时")
			touText := imgNames[i] + "连接超时"
			upload.WriteFile(touText)
			continue
		}
		total++
	}
	totalText := "测试完成"
	fmt.Println(totalText)
	upload.WriteFile(totalText)
}

func main() {
	cleanCycle("https@api-intl%seeunsee%cn+v1+identify")

}
