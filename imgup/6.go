package main

import (
	"net/http"
	"bytes"
	"fmt"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
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

		//fmt.Println()
	}
	return body.String()
}
func main() {



	a:=doUpload("http://m.seeunsee.cn/intelligent-packaging-check/check.php","C:/Users/huash/Documents/GoPlace/src/GoPlace/imgup/image/cigarette.jpg")
	fmt.Println(a)
}
