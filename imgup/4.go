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

func doUpload(url string, path string) (body bytes.Buffer) {
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
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()

		//fmt.Println(body.String())

	}
	return body
}

func inputFile(num int) (url string, result string) {
	sec := strconv.Itoa(num)
	inputFilePath, _ := os.Getwd()
	inputFilePath += "/input.ini"
	cfg, err := ini.Load(inputFilePath)
	if err != nil {
		fmt.Println(err)
	}
	url, err := cfg.Section(sec).GetKey("url")
	if err != nil {
		fmt.Println(err)
	}
	result, err := cfg.Section(sec).GetKey("result")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
	fmt.Println(result)
	return url, result
}

func getImgPath(num int) (imgPath string) {
	path, _ := os.Getwd()
	nums := strconv.Itoa(num)
	imgPath = path + "/image/" + nums + ".jpg"
	return imgPath
}

func judge(result string) (judgement bool) {

}

func outputFile(url string, num int, judgement bool) {

}

func main() {

}
