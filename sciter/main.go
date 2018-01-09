package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"upload"
	"fmt"
	"os"
	"strconv"
)

func check(err error)  {
	if err!=nil{
		fmt.Println(err)
	}
}

//获取csv文件夹中的文件名称转换为api地址，加入到下拉框选项中
func addUrlOption(root *sciter.Element){
	for num:=0;num<len(upload.CsvNameSlice());num++{
		url:=upload.ApiUrlSlice(upload.CsvNameSlice())[num]
		add,_:=sciter.CreateElement("option",url)
		add.SetAttr("value",upload.CsvNameSlice()[num])
		add.SetAttr("url",url)
		a,_:=root.SelectById("#url")
		err:=a.Insert(add,num+1)
		check(err)
	}
}

//获取选择文件夹的图片数量
func defineFunc(w *window.Window){
	w.DefineFunction("picNum", func(args ...*sciter.Value) *sciter.Value {
		_,_,picNum,_,_:=upload.ReadCsvFile(args[0].String())
		return sciter.NewValue(picNum)
	})
}

//将图片数量转换成文字
func defMethod(root *sciter.Element) {
	numP, _ := root.SelectById("imgNum")
	numP.DefineMethod("picPa", func(args ...*sciter.Value) *sciter.Value {
		numText:="共有"+args[0].String()+"张图片"
		numP.SetText(numText)
		return sciter.NullValue()
	})
}

//创建li元素
func printLi(addText string,obj *sciter.Element)  *sciter.Value{
	add,_:=sciter.CreateElement("li",addText)
	//add.SetAttr("class","output")
	err:=obj.Insert(add,1)
	check(err)
	return sciter.NullValue()
}


func defMethod1(root *sciter.Element,w *window.Window)  {
	//上传按钮
	beginBtn,_:=root.SelectById("begin")
	//args:csvName
	beginBtn.DefineMethod("startUpload", func(args ...*sciter.Value) *sciter.Value {
		upload.WriteTimeToFile()
		url:=upload.ApiUrlText(args[0].String())
		upload.WriteFile(url)
		imgName, rightResult, _,paramName,extraParam:=upload.ReadCsvFile(args[0].String())
		//把图片名称转换为图片地址
		currentPath,_:=os.Getwd()
		var imgPath []string
		for _,k:=range imgName{
			absPath := currentPath + "/image/" + args[0].String() + "/" +k
			imgPath=append(imgPath,absPath)
		}
		//开始上传并输出
		var respBody string
		var total,sucs,fail=0,0,0
		//outPutObj,_:=root.SelectById("#output")
		for i, j :=range imgPath{

			respBody=upload.DoUpload(url, j,paramName,extraParam)
			fmt.Println(respBody)
			total++
			fmt.Println("1")
			if upload.Imatch(rightResult[i],respBody){
				fmt.Println("2")
				sucs++
				succText:=imgName[i]+"返回结果："+respBody+"，识别成功。"
				upload.WriteFile(succText)
				//printLi(succText,outPutObj)
				beginBtn.CallMethod("fac",sciter.NewValue(succText))
			}else {
				fmt.Println("3")
				fail++
				failText:=imgName[i]+"返回结果："+respBody+"，正确结果："+rightResult[i]+"，识别失败。"
				upload.WriteFile(failText)
				//printLi(failText,outPutObj)
				beginBtn.CallMethod("fac",sciter.NewValue(failText))
			}
		}
		fmt.Println("4")
		sumObj,_:=root.SelectById("summery")
		totalText:="共上传"+strconv.Itoa(total)+"张图片，其中识别成功"+strconv.Itoa(sucs)+"张，识别失败"+strconv.Itoa(fail)+"张。"
		upload.WriteFile(totalText)
		printLi(totalText,sumObj)
		return sciter.NullValue()
	})
}

func main() {
	//创建窗口
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG,
		//设置窗口大小
		&sciter.Rect{Left: 300, Top: 300, Right: 800, Bottom: 800})
	check(err)
	//载入html
	w.LoadFile("F:/Code/GoPlace/sciter/demo.html")

	root, _ := w.GetRootElement()
	addUrlOption(root)
	defineFunc(w)
	defMethod(root)
	defMethod1(root,w)


	//设置标题
	w.SetTitle("图片上传接口测试程序")
	//显示窗口
	w.Show()
	//运行
	w.Run()
}
