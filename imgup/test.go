package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"os"
	"net/http"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
	"strconv"
	"github.com/go-ini/ini"
)

func getIniFileName(path string) ([]string){
	//输入ini文件夹路径，输出包含所有ini文件名称的slice
	f,err:=ioutil.ReadDir(path)
	if err!=nil{
		fmt.Println(err)
	}
	var iniFile []string
	for _, n := range f {
		v:=n.Name()
		reg,_:=regexp.MatchString(".ini",v)
		if reg{
			iniFile =append(iniFile,v)
		}
	}
	return iniFile
}

func getUrl(path string,num int)(url string ,url1 string){
	//输入ini文件夹路径，输出ini文件名所代表的接口url
	url1=strings.Replace(getIniFileName(path)[num],".ini","",-1)
	url2:=strings.Replace(url1,"+","/",-1)
	url3:=strings.Replace(url2,"%",".",-1)
	url="http://"+url3
	return url,url1
}

func check(err error){
	if err!=nil{
		fmt.Println(err)
	}
}

func newRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	check(err)
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	check(err)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func doUpload(url string, path string) string {
	body := &bytes.Buffer{}
	extraParams := map[string]string{
		"type": "1/jpeg",
	}
	req, err := newRequest(url, extraParams, "file", path)
	check(err)
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
	}
	return body.String()
}

func analysisIni(iniPath string,num int) (result string) {
	sec := strconv.Itoa(num)
	cfg, err := ini.Load(iniPath)
	check(err)
	result = cfg.Section(sec).Key("result").String()
	return result
}

func match(result string, respBody string) bool {
	rt := regexp.QuoteMeta(result)
	ry := regexp.QuoteMeta(respBody)
	return strings.Contains(ry, rt)
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func main(){
	filePath,_:=os.Getwd()
	//ini文件夹路径
	iniFilePath :=filePath+"/ini/"
	iniFileList:=getIniFileName(iniFilePath)
	//创建log.txt
	//f,err:=os.OpenFile("log.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	//check(err)
	fmt.Println("1")

	for num:=1;num>=len(iniFileList);num++{
		fmt.Println("2")
		//http://dshfjsbadhfbs.dsf+图片文件夹名称
		url,imgPathName:=getUrl(iniFilePath,num)
		fmt.Println("3")
		//ini文件路径
		iniPath:=iniFilePath+iniFileList[num]
		//图片文件夹名称
		//_,imgPathName:= getUrl(iniFilePath,num)

		for numb:=1;;numb++{
			imgPath:=filePath+"/image/"+imgPathName+"/"+strconv.Itoa(numb)
			if pathExist(imgPath){
				//上传
				respBody := doUpload(url, imgPath)

				result:=analysisIni(iniPath,numb)

				judge:=match(result,respBody)
				if judge{
					fmt.Println("3")
				}else {
					fmt.Println("4")
				}
			}else {
				break
			}
		}
	}
}
