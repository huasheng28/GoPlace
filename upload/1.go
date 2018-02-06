package main

import (
	"os"
	"bytes"
	"mime/multipart"
	"io"
	"fmt"
	"time"
	"net"
	"net/http"

	"net/textproto"

	"path/filepath"
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}
//var quoteEscaper = strings.NewReplacer("//", "////", `"`, "///"")
//
//func escapeQuotes(s string) string {
//	return quoteEscaper.Replace(s)
//}
//func CreatehhFormFile(w *multipart.Writer,fieldname, filename string) (io.Writer, error) {
//	h := make(textproto.MIMEHeader)
//	h.Set("Content-Disposition",
//		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
//			escapeQuotes(fieldname), escapeQuotes(filename)))
//	h.Set("Content-Type", "application/octet-stream")
//	return w.CreatePart(h)
//}
func CreateExFormFile(w *multipart.Writer, fieldname string,filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,fieldname, filename)) //+";type=image/jpeg"
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
	if headers!=nil{
		fileName=filepath.Base(path)+";type=image/jpeg"
	}else {
		fileName=filepath.Base(path)
	}
	part, err := CreateExFormFile(writer,paramName,fileName)
	//part, err := writer.CreateFormFile("pic", filepath.Base(path))
	check(err)

	_, err = io.Copy(part, file)
	check(err)

	//a,_:=writer.CreateFormField("Content-Type")
	//a.Write([]byte("image/jpeg"))
	if extraParam!=nil{
		for key, val := range extraParam {
			_ = writer.WriteField(key, val)
		}
	}
	//contentType := writer.FormDataContentType()
	//fmt.Println(contentType)
	writer.Close()

	req, err := http.NewRequest("POST", url, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if headers!=nil{
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	fmt.Println(req.Header)
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
		bodyz = ""
	} else if err != nil {
		fmt.Println(err)
	} else {
		_, err := body.ReadFrom(resp.Body)
		check(err)
		defer resp.Body.Close()
		bodyz = body.String()
	}
	fmt.Println(bodyz)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	return bodyz
}

func main() {
	DoUpload("https://seeunsee.cn/api/tx-check-v240.php","F:/Code/GoPlace/upload/image/https@seeunsee%cn+api+tx-check-v240%php/1.jpg",
		"pic",map[string]string{"access_token":"793970d616c5b8281ed89c41d1c8f5cb",},nil)
	//DoUpload("http://192.168.1.178/recognize3.php","F:/Code/GoPlace/upload/image/http@192%168%1%178+recognize3%php/1.jpg","file",
	//	nil,map[string]string{"X-Api-Key":"CFE95B64AC715D64275365EDE690GH7C",})
}