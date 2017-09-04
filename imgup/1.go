package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
func upfile() {
	path, _ := os.Getwd()
	path += "/pian.jpg"
	extraParams := map[string]string{
		"id":               "WU_FILE_0",
		"name":             "pian.jpg",
		"type":             "image/jpeg",
		"lastModifiedDate": "Mon Apr 17 2017 15:31:08 GMT+0800 (CST)",
		"size":             "241281",
	}
	_, err := newfileUploadRequest("http://test.seeunsee.cn/jsb/wm-test/api/check.php", extraParams, "file", path)
	if err != nil {
		log.Fatal(err)
	}
}

func up() {
	for a := 1; a < 3; a++ {
		go upfile()
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(body)
	}
}
func main() {
	for a := 0; a < 3; a++ {
		up()
	}
	for {

	}
}
