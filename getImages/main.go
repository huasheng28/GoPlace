package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
	"strings"
	"strconv"
)

func check(err error) {
	if err != nil{
		fmt.Println(err)
	}
}

func readCsv() [][]string {
	dir, err := os.Getwd()
	check(err)
	csvFile := dir + "/lost.csv"
	check(err)
	file, err := os.Open(csvFile)
	check(err)
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)
	return records
}

func getImages(url string) {
	name := strings.Split(url, "/")
	fmt.Print(name[len(name)-1])
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name[len(name)-1])
	io.Copy(out, bytes.NewReader(body))
}

func main() {
	for i,record :=range readCsv(){
		getImages(record[0])
		j := strconv.Itoa(i+1)
		fmt.Println(" 第 "+j+" 张完成")
	}
	fmt.Println("全部完成")
}
