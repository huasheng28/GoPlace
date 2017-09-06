package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(url, fileName, filePath string) error {

	//打开并读取文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer file.Close()

	//创建form
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//设置文件的上传参数叫uploadfile, 文件名是filename,相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("upload", fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println("err1")
		return err
	}

	//获取上传文件的类型
	contentType := bodyWriter.FormDataContentType()
	//contentType:=
	bodyWriter.Close()

	//上传的其他参数设置
	params := map[string]string{

		"file": fileName,
		//"shap":         shap,
		//"access_token": access_token,
		//"channel":      channel,
	}
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}

	//发送post请求
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		fmt.Println("err2")
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err3")
		return err
	}
	fmt.Println("status is", resp.Status)
	fmt.Println("result is", string(respBody))

	return nil
}

//func readFile() {
//
//}
func main() {
	url := "http://m.seeunsee.cn/intelligent-packaging-check/check.php"
	fileName := "cigarette.jpg"
	filePath, _ := os.Getwd()
	filePath += "/image/cigarette.jpg"
	//shap := "1"
	//access_token := "793970d616c5b8281ed89c41d1c8f5cb"
	//channel := "1"
	postFile(url, fileName, filePath)
}
