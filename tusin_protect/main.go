package main

import (
	"tusin_protect"
	"fmt"
)

//打包两个string类型切片，一一对应生成二维切片
func zip(a, b []string) ([][]string, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}
	r := make([][]string, len(a), len(a))
	for i, e := range a {
		r[i] = []string{e, b[i]}
	}
	return r, nil
}

func main()  {
	delaytime,outtime,urlSlice,pathSlice,phoneSlice,emailSlice,receiverSlice:=tusin_protect.ReadCsv()
	a,err:=zip(urlSlice,pathSlice)
	if err!=nil{
		fmt.Println(err)
	}
	for i:=0;i<len(a);i++{
		respbody:=tusin_protect.Urltest(a[i][0],outtime)
		if respbody{
			continue
		}else {
			tusin_protect.Restarts(a[i][1])
			tusin_protect.SendEmail(emailSlice[0],emailSlice[1],receiverSlice)
		}
	}

}