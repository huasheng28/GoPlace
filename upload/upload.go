package upload

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"encoding/csv"
	"regexp"
	"encoding/json"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//上传文件操作
func newRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	check(err)
	_, err = io.Copy(part, file)
	check(err)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	check(err)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
func DoUpload(url string, path string, paramName string,extraParam map[string]string) string {
	body := &bytes.Buffer{}
	req, err := newRequest(url, extraParam, paramName, path)
	check(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		_, err := body.ReadFrom(resp.Body)
		check(err)
		resp.Body.Close()
	}
	return body.String()
}

//获取csv文件名切片，不带后缀
func CsvNameSlice() []string {
	currentPath ,_:= os.Getwd()
	csvFilePath := currentPath + "/csv/"
	f, err := ioutil.ReadDir(csvFilePath)
	check(err)
	var csvNameSlice []string
	for _, n := range f {
		v := n.Name()
		reg, _ := regexp.MatchString(".csv", v)
		if reg {
			b := strings.Replace(v, ".csv", "", -1)
			csvNameSlice = append(csvNameSlice, b)
		}
	}
	return csvNameSlice
}

//将csv文件名转换成api地址
func ApiUrlSlice(csvNameSlice []string) []string {
	var apiUrlSlice []string
	for _, a := range csvNameSlice {
		b := strings.Replace(a, "+", "/", -1)
		c := strings.Replace(b, "%", ".", -1)
		d := "http://" + c
		apiUrlSlice = append(apiUrlSlice, d)
	}
	return apiUrlSlice
}

func ApiUrlText(csvName string) string {
	b := strings.Replace(csvName, "+", "/", -1)
	c := strings.Replace(b, "%", ".", -1)
	d := "http://" + c
	return d
}

//读取csv文件，返回文件中的参数和参数数量
func ReadCsvFile(csvName string) ([]string, []string, int,string,map[string]string) {
	//读取文件
	currentPath, _ := os.Getwd()
	csvFile := currentPath + "/csv/" + csvName + ".csv"
	f, err := os.Open(csvFile)
	check(err)
	defer f.Close()
	r := csv.NewReader(f)
	record, err := r.ReadAll()
	check(err)
	//图片数量
	picNum := len(record)-1
	//上传参数
	paramName:=record[0][0]
	//其他参数
	var extraParam map[string]string
	extraParamJson:=record[0][1]
	if extraParamJson==""{
		extraParam=map[string]string{}
	}else {
		err1 := json.Unmarshal([]byte(extraParamJson), &extraParam)
		check(err1)
	}
	//图片地址及正确返回参数
	var imgName, rightResult []string
	for _, a := range record[1:] {
		imgName = append(imgName, a[0])
		rightResult = append(rightResult, a[1])
	}
	return imgName, rightResult, picNum,paramName,extraParam
}

//判断返回数据是否与cav文件中一致
func Imatch(rightResult string,respBody string) bool{
	a:=regexp.QuoteMeta(rightResult)
	b:=regexp.QuoteMeta(respBody)
	return strings.Contains(b, a)
}

//输出文件
func WriteFile(outText string){
	f, err := os.OpenFile("log.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	f.WriteString(outText+",\r\n")
}
//输出当前时间到文件
func WriteTimeToFile()  {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	text:=tm.Format("2006-01-02 03:04:05 PM")
	WriteFile(text)
}