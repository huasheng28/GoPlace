package tusin_protect

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Urltest(url string, outtime int) (respbody bool) {
	timeout := time.Duration(outtime) * time.Second
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add()
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if e, ok := err.(net.Error); ok && e.Timeout() {
		respbody = false
		fmt.Println("超时")
	} else if err != nil {
		fmt.Println(err)
		return false
	} else {
		code := resp.StatusCode
		if code == 200 {
			fmt.Println("返回成功")
			return true
		}
	}
	defer resp.Body.Close()
	return
}
