package imgUpload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"time"
)

type ARequest struct {
	Url       string            //网络地址
	Headers   map[string]string //请求头
	Params    map[string]string //body变量
	ImgParam  string            //上传文件对应的变量名
	ImgFolder string            //图像文件夹地址
	OutTime   int               //超时时间
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		WriteToFile("Error:"+err.Error())
	}
}

func createFormFile(w *multipart.Writer, fieldname string, filename string) (io.Writer, error) {
	//构成请求体格式
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filename))
	h.Set("Content-Type", "image/jpeg")
	return w.CreatePart(h)
}

func newRequest(reqHand ARequest, imgPath string) (*http.Request, error) {
	//读取图片
	file, err := os.Open(imgPath)
	check(err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	//判断是否存在请求头
	//后缀是否添加有待商榷，加了有时候会出问题
	var fileName string
	//if len(reqHand.Headers) != 0 {
	//	fileName = filepath.Base(reqHand.ImgFolder) + ";type=image/jpeg"
	//} else {
	fileName = filepath.Base(imgPath)
	//}

	//添加请求体格式
	part, err := createFormFile(writer, reqHand.ImgParam, fileName)
	check(err)
	_, err = io.Copy(part, file)
	check(err)
	if reqHand.Params != nil {
		for k, v := range reqHand.Params {
			writer.WriteField(k, v)
		}
	}

	req, err := http.NewRequest("POST", reqHand.Url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	//req.Header.Add("ContentLength",strconv.Itoa(body.Len()+68))
	if reqHand.Headers != nil {
		for k, v := range reqHand.Headers {
			req.Header.Add(k, v)
		}
	}

	//为什么多了68字节啊？？？？？？
	req.ContentLength=int64(body.Len()+68)
	return req, err
}

func DoUpload(reqHand ARequest, imgPath string) (respBody string) {
	bodya := &bytes.Buffer{}
	req, err := newRequest(reqHand, imgPath)
	check(err)
	client := http.Client{
		Timeout: time.Duration(reqHand.OutTime) * time.Second,
	}

	resp, err := client.Do(req)

	if e, ok := err.(net.Error); ok && e.Timeout() {
		respBody = "timeout"
	} else if err != nil {
		check(err)
	} else {
		_, err := bodya.ReadFrom(resp.Body)
		check(err)
		defer resp.Body.Close()
		respBody = bodya.String()
	}
	fmt.Println(respBody)
	return respBody
}
