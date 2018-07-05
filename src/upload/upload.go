package upload

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//上传文件操作
func CreateExFormFile(w *multipart.Writer, fieldname string, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filename))
	h.Set("Content-Type", "image/jpeg")
	return w.CreatePart(h)
}

var fileName string

func newRequest(url string, extraParam map[string]string, paramName, path string, headers map[string]string) (*http.Request, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//if len(headers) != 0 {
	//	fileName = filepath.Base(path) + ";type=image/jpeg"
	//} else {
	fileName = filepath.Base(path)
	//}
	fmt.Println(fileName)
	part, err := CreateExFormFile(writer, paramName, fileName)
	//part, err := writer.CreateFormFile("pic", filepath.Base(path))
	check(err)
	_, err = io.Copy(part, file)
	check(err)
	if extraParam != nil {
		for key, val := range extraParam {
			_ = writer.WriteField(key, val)
		}
	}
	writer.Close()
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	return req, err
}

func DoUpload(url string, path string, paramName string, extraParam map[string]string, headers map[string]string) (bodyz string) {
	body := &bytes.Buffer{}
	req, err := newRequest(url, extraParam, paramName, path, headers)
	check(err)
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Do(req)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		bodyz = "timeout"
	} else if err != nil {
		fmt.Println(err)
	} else {
		_, err := body.ReadFrom(resp.Body)
		check(err)
		defer resp.Body.Close()
		bodyz = body.String()
	}
	fmt.Println(bodyz)
	return bodyz
}

//获取csv文件名切片，不带后缀
func CsvNameSlice() []string {
	currentPath, _ := os.Getwd()
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
		d := strings.Replace(c, "@", "://", -1)
		apiUrlSlice = append(apiUrlSlice, d)
	}
	return apiUrlSlice
}

func ApiUrlText(csvName string) string {
	b := strings.Replace(csvName, "+", "/", -1)
	c := strings.Replace(b, "%", ".", -1)
	d := strings.Replace(c, "@", "://", -1)
	return d
}

//读取csv文件，返回文件中的参数和参数数量
func ReadCsvFile(csvName string) ([]string, []string, int, string, map[string]string, map[string]string) {
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
	picNum := len(record) - 1
	//上传参数
	paramName := record[0][0]
	var extraParam, headers map[string]string
	//param
	if len(record[0]) == 2 {
		extraParamJson := record[0][1]
		if extraParamJson == "" {
			extraParam = nil
		} else {
			err1 := json.Unmarshal([]byte(extraParamJson), &extraParam)
			check(err1)
		}
		headers = map[string]string{}
	} else if len(record[0]) == 3 {
		extraParamJson := record[0][1]
		if extraParamJson == "" {
			extraParam = nil
		} else {
			err1 := json.Unmarshal([]byte(extraParamJson), &extraParam)
			check(err1)
		}
		//请求头
		headersJson := record[0][2]
		if headersJson == "" {
			headers = nil
		} else {
			err2 := json.Unmarshal([]byte(headersJson), &headers)
			check(err2)
		}
	} else {
		panic("wrong file format")
	}
	//图片地址及正确返回参数
	var imgName, rightResult []string
	for _, a := range record[1:] {
		imgName = append(imgName, a[0])
		rightResult = append(rightResult, a[1])
	}
	return imgName, rightResult, picNum, paramName, extraParam, headers
}

//判断返回数据是否与cav文件中一致
func Imatch(rightResult string, respBody string) bool {
	a := regexp.QuoteMeta(rightResult)
	b := regexp.QuoteMeta(respBody)
	return strings.Contains(b, a)
}

//输出文件
func WriteFile(outText string) {
	f, err := os.OpenFile("log.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	enc := mahonia.NewEncoder("gbk").ConvertString(outText)
	f.WriteString(enc + ",\r\n")
	return
}

//输出当前时间到文件
func WriteTimeToFile() {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	text := tm.Format("2006-01-02 03:04:05 PM")
	WriteFile(text)
}

//将返回值中的逗号替换掉
func RespCsv(respBody string) string {
	return strings.Replace(respBody, ",", "%", -1)
}
