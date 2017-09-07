package main

import (
	"os"
	"fmt"
	"net/http"
	"bytes"
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
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println(body.String())

	}
	return body
}
func main()  {

}
