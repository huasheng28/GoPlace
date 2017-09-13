package main

import (
	"bytes"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func getIniFileName(path string) []string {
	//输入ini文件夹路径，输出包含所有ini文件名称的slice
	f, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	var iniFile []string
	for _, n := range f {
		v := n.Name()
		reg, _ := regexp.MatchString(".ini", v)
		if reg {
			iniFile = append(iniFile, v)
		}
	}
	return iniFile
}

func getUrl1(path string, num int) (url1 string) {
	//输入ini文件夹路径，输出img文件夹名称
	return strings.Replace(getIniFileName(path)[num], ".ini", "", -1)
}

func getUrl(url1 string) string {
	//输出url
	url2 := strings.Replace(url1, "+", "/", -1)
	url3 := strings.Replace(url2, "%", ".", -1)
	url := "http://" + url3
	return url
}

func check(err error) {
	if err != nil {
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

func analysisIni(iniPath string, num int) (result string) {
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

func main() {
	filePath, _ := os.Getwd()
	//ini文件夹路径
	iniFilePath := filePath + "/ini/"
	iniFileList := getIniFileName(iniFilePath)

	//创建log.txt
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	check(err)
	i, j, k := 0, 0, 0

	for num := 0; num < len(iniFileList); num++ {
		//图片文件夹名称
		imgPathName := getUrl1(iniFilePath, num)
		//http://dshfjsbadhfbs.dsf
		url := getUrl(imgPathName)
		fmt.Println(url)
		//ini文件路径
		iniPath := iniFilePath + iniFileList[num]
		f.WriteString(url + "\r\n")

		for numb := 1; ; numb++ {
			imgPath := filePath + "/image/" + imgPathName + "/" + strconv.Itoa(numb) + ".jpg"
			if pathExist(imgPath) {
				i++
				f.WriteString("第" + strconv.Itoa(i) + "次上传\r\n")

				//上传
				respBody := doUpload(url, imgPath)
				fmt.Println(respBody)

				result := analysisIni(iniPath, numb)

				f.WriteString(respBody + "\r\n")
				judge := match(result, respBody)
				if judge {
					f.WriteString("匹配成功\r\n")
					fmt.Println("成功")
					j++
				} else {
					f.WriteString("匹配失败\r\n")
					fmt.Println("失败")
					k++
				}
			} else {
				break
			}
		}
	}
	f.WriteString("共上传" + strconv.Itoa(i) + "次 ")
	f.WriteString("成功" + strconv.Itoa(j) + "次 ")
	f.WriteString("失败" + strconv.Itoa(k) + "次\r\n")
	defer f.Close()
}
