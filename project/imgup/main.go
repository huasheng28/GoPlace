package main

import (
	"bufio"
	"fmt"
	"github.com/huasheng28/GoPlace/upload"
	"os"
	"strconv"
)

func startCycle(csvName string) {
	upload.WriteTimeToFile()
	url := upload.ApiUrlText(csvName)
	upload.WriteFile(url)
	imgName, rightResult, _, paramName, extraParam := upload.ReadCsvFile(csvName)
	//把图片名称转换为图片地址
	currentPath, _ := os.Getwd()
	var imgPath []string
	for _, k := range imgName {
		absPath := currentPath + "/image/" + csvName + "/" + k
		imgPath = append(imgPath, absPath)
	}
	//开始上传并输出
	var respBody,respCsv,rightCsv string
	var total, sucs, fail = 0, 0, 0
	for i, j := range imgPath {
		respBody = upload.DoUpload(url, j, paramName, extraParam)
		respCsv =upload.RespCsv(respBody)
		total++
		if upload.Imatch(rightResult[i], respBody) {
			sucs++
			succText := imgName[i] + "返回结果：" + respCsv + "，识别成功。"
			fmt.Println(succText)
			upload.WriteFile(succText)
		} else {
			fail++
			rightCsv=upload.RespCsv(rightResult[i])
			failText := imgName[i] + "返回结果：" + respCsv + "，正确结果：" + rightCsv + "，识别失败。"
			fmt.Println(failText)
			upload.WriteFile(failText)
		}
		fmt.Println("共" + strconv.Itoa(total) + "张图片，其中识别成功" + strconv.Itoa(sucs) + "张，识别失败" + strconv.Itoa(fail) + "张。")
	}
	totalText := "共" + strconv.Itoa(total) + "张图片，其中识别成功" + strconv.Itoa(sucs) + "张，识别失败" + strconv.Itoa(fail) + "张。"
	upload.WriteFile(totalText)
}

func main() {
	//显示url地址
	for num := 0; num < len(upload.CsvNameSlice()); num++ {
		url := upload.ApiUrlSlice(upload.CsvNameSlice())[num]
		fmt.Println(num)
		fmt.Println(url)
	}
	fmt.Println("请输入需要检验的接口编号：")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		index, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("数据转换错误", err)
		}
		startCycle(upload.CsvNameSlice()[index])
		fmt.Println("请输入需要检验的接口编号：")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
