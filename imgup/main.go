package main

import (
	"bufio"
	"fmt"
	"upload"
	"os"
	"strconv"
)

func startCycle(csvName string) {
	upload.WriteTimeToFile()
	url := upload.ApiUrlText(csvName)
	upload.WriteFile(url)
	imgName, rightResult, picNum, paramName, extraParam := upload.ReadCsvFile(csvName)
	//把图片名称转换为图片地址
	currentPath, _ := os.Getwd()
	var imgPath []string
	for _, k := range imgName {
		absPath := currentPath + "/image/" + csvName + "/" + k
		imgPath = append(imgPath, absPath)
	}
	//开始上传并输出
	var respBody, respCsv, rightCsv string
	var total, sucs, fail, timeout = 0, 0, 0, 0
	for i, j := range imgPath {
		fmt.Printf("正在测试[" + strconv.Itoa(i+1) + "/" + strconv.Itoa(picNum) + "]...")
		respBody = upload.DoUpload(url, j, paramName, extraParam)
		if respBody == "" {
			timeout++
			fmt.Println("连接超时")
			touText := imgName[i] + "连接超时"
			upload.WriteFile(touText)
			continue
		}
		respCsv = upload.RespCsv(respBody)
		total++
		if upload.Imatch(rightResult[i], respBody) {
			sucs++
			succText := imgName[i] + "，返回结果：" + respCsv + "，识别成功。"
			fmt.Println("成功")
			upload.WriteFile(succText)
		} else {
			fail++
			rightCsv = upload.RespCsv(rightResult[i])
			failText := imgName[i] + "，返回结果：" + respCsv + "，正确结果：" + rightCsv + "，识别失败。"
			fmt.Println("失败")
			upload.WriteFile(failText)
		}
	}
	totalText := "测试完成，共" + strconv.Itoa(sucs) + "成功" + strconv.Itoa(fail) + "失败" + strconv.Itoa(timeout) + "超时"
	fmt.Println(totalText)
	upload.WriteFile(totalText)
}

//判断数值是否属于切片
func contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func main() {
	//显示url地址
	csvNames := upload.CsvNameSlice()
	var numLimit []int
	for num := 0; num < len(csvNames); num++ {
		url := upload.ApiUrlSlice(csvNames)[num]
		numLimit = append(numLimit, num)
		fmt.Printf(strconv.Itoa(num) + "--")
		fmt.Println(url)
	}

	fmt.Println("请输入需要检验的接口编号：")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		if index, err := strconv.Atoi(line); err != nil {
			fmt.Println("输入值类型错误，请重新输入：")
			continue
		} else if !contains(numLimit, index) {
			fmt.Println("输入值超出范围，请重新输入：")
			continue
		}
		index, _ := strconv.Atoi(line)
		startCycle(csvNames[index])
		for num := 0; num < len(csvNames); num++ {
			url := upload.ApiUrlSlice(csvNames)[num]
			fmt.Printf(strconv.Itoa(num) + "--")
			fmt.Println(url)
		}
		fmt.Println("继续其他接口测试：")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
