package main

import (
	"bytes"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"regexp"
)

func newRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func doUpload(url string, path string) string {
	body := &bytes.Buffer{}
	extraParams := map[string]string{
		"type": "image/jpeg",
	}
	req, err := newRequest(url, extraParams, "file", path)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {

		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()

		//fmt.Println(body.String())

	}
	return body.String()
}

func inputFile(num int) (url string, result string) {
	sec := strconv.Itoa(num)
	inputFilePath, _ := os.Getwd()
	inputFilePath += "/input.ini"
	cfg, err := ini.Load(inputFilePath)
	if err != nil {
		fmt.Println(err)
	}
	url = cfg.Section(sec).Key("url").String()
	result = cfg.Section(sec).Key("result").String()
	return url, result
}

func getImgPath(num int) (imgPath string) {
	path, _ := os.Getwd()
	nums := strconv.Itoa(num)
	imgPath = path + "/image/" + nums + ".jpg"
	return imgPath
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func match(result string,respBody string) (bool,error){
	return regexp.MatchString(result,respBody)
}

func outputFile(url string, num int, judgement bool) {

}

func main() {
	for num := 1; ; num++ {
		//取图片路径
		imgPath := getImgPath(num)

		//判断路径是否存在
		if pathExist(imgPath) {

			//取ini文件数据
			url, result := inputFile(num)

			//上传图片取返回参数
			respBody := doUpload(url, imgPath)

			//判断结果是否正确
			judge,err:=match(result,respBody)
			if err!=nil{
				fmt.Println(err)
			}
			if judge{
				fmt.Println("1")
			}else {
				fmt.Println("2")
			}

		} else {
			break
		}
	}
}
