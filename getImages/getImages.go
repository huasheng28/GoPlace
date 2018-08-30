package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
)

//func check(err error) {
//	if err != nil{
//		fmt.Println(err)
//		f, _ := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//		defer f.Close()
//		_, err = f.WriteString(err.Error()+"\r\n")
//	}
//}

//func readCsv() [][]string {
//	dir, err := os.Getwd()
//	check(err)
//	csvFile := dir + "/lost.csv"
//	check(err)
//	file, err := os.Open(csvFile)
//	check(err)
//	defer file.Close()
//	reader := csv.NewReader(file)
//	records, err := reader.ReadAll()
//	check(err)
//	return records
//}

func getImages(url1 string)[]byte{
	fmt.Println("http://"+url1[10:])
	a:="http://"+url1[10:]
	resp, err := http.Get(a)
	check(err)
	fmt.Println("status: ",resp.Status)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body
}

func writeImg(body []byte,name string,path string) {
	path=getFolder(path)
	fullName :=path+"\\"+name+".JPG"
	fmt.Println(fullName)
	out, err := os.Create(fullName)
	check(err)
	io.Copy(out, bytes.NewReader(body))
}

func getFolder(num string)string{
	dir, err := os.Getwd()
	check(err)
	path:=dir+"\\"+num
	if _, err := os.Stat(path); err != nil {
		fmt.Println("创建目录 ", path)
		err := os.MkdirAll(path, 0711)
		if err != nil {
			fmt.Println("目录创建出错")
			fmt.Println(err)
		}
	}
	return path
}

func main() {
	//var path string
	//for _,record :=range readCsv(){
	//	path=getFolder(record[1])
	//	getImages(record[0],record[2],path)
	//	fmt.Println(" 第 "+record[2]+" 张完成")
	//}
	//fmt.Println("全部完成")
	dir, err := os.Getwd()
	check(err)
	csvFile := dir + "/lost.csv"
	check(err)
	file, err := os.Open(csvFile)
	check(err)
	defer file.Close()
	reader := csv.NewReader(file)
	for{
		record,err:=reader.Read()
		if err==io.EOF{
			break
		}else if err!=nil {
			fmt.Println(err)
			return
		}
		body1:=getImages(record[0])
		writeImg(body1,record[2],record[1])
		fmt.Println(" 第 "+record[2]+" 张完成")
	}
	fmt.Println("全部完成")
}
