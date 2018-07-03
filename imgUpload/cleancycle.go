package main

import (
	"imgUpload"
	"fmt"
	"strconv"
	"path/filepath"
	"regexp"
)

func main() {
	reqSlice := imgUpload.ReadEntryCsv()
	var text,respBody,imgName string
	for _,req:=range reqSlice{
		imgPathSlice := imgUpload.ImgPathSlice(req.ImgFolder)
		for i,imgPath :=range imgPathSlice{
			fmt.Printf("正在测试[" + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(imgPathSlice)) + "]...")
			fmt.Println(imgPath)
			imgName = filepath.Base(imgPath)
			respBody = imgUpload.DoUpload(req, imgPath)

			if respBody=="timeout"{
				text = imgName + "连接超时重试一次"
				fmt.Println(text)
				imgUpload.WriteToFile(text)
				respBody = imgUpload.DoUpload(req, imgPath)
			}else if ok, err :=regexp.MatchString("\"errcode\":50001",respBody);ok{
				if err !=nil {
					imgUpload.WriteToFile(err.Error())
				}
				text = imgName+"返回错误重试一次"
				fmt.Println(text)
				imgUpload.WriteToFile(text)
				respBody = imgUpload.DoUpload(req, imgPath)
			}
			imgUpload.WriteTime()
			imgUpload.WriteToFile(imgName)
			imgUpload.WriteToFile(respBody)
		}
	}
	text = "测试完成"
	fmt.Println(text)
	imgUpload.WriteTime()
	imgUpload.WriteToFile(text)
}
