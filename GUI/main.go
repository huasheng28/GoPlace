package upload

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
)

func defFunc(root *sciter.Element) {
	uploadBtn, err := root.SelectById("uploadImg")
	if err != nil {
		log.Fatal(err)
	}
	uploadBtn.DefineMethod("upload", func(args ...*sciter.Value) *sciter.Value {
		//url:=args[0].String()
		//path:=args[1].String()
		//doUpload(url,path)
		return sciter.NullValue()
	})
}

func main() {
	//创建窗口
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG,
		//设置窗口大小
		&sciter.Rect{Left: 300, Top: 300, Right: 800, Bottom: 800})
	if err != nil {
		log.Fatal(err)
	}

	//载入html
	w.LoadFile("F:/Code/GoPlace/GUI/demo.html")

	root, _ := w.GetRootElement()
	defFunc(root)

	//设置标题
	w.SetTitle("图片上传接口测试程序")
	//显示窗口
	w.Show()
	//运行
	w.Run()
}
