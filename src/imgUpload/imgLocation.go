package imgUpload

import "io/ioutil"

func TraversFolder(imgFolder string) (imgNames []string) {
	//遍历文件夹，获取所有文件名
	files, _ := ioutil.ReadDir(imgFolder)
	for _,file := range files{
		imgNames = append(imgNames, file.Name())
	}
	return imgNames
}

func ImgPathSlice(imgFolder string) (imgPathSlice []string) {
	imgNameSlice :=TraversFolder(imgFolder)
	var imgPath string
	for _,imgName :=range imgNameSlice{
		imgPath = imgFolder+"/"+imgName
		imgPathSlice =append(imgPathSlice,imgPath)
	}
	return imgPathSlice
}