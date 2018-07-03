package imgUpload

import (
	"os"
	"encoding/csv"
	"encoding/json"
	"strconv"
)


func GetRootFolder() string {
	//取当前目录地址
	dir, err := os.Getwd()
	check(err)
	return dir
}

func stringToMap(recordString string) (jsonMap map[string]string) {
	//将json格式转换成map
	if recordString == "" {
		return nil
	}
	err := json.Unmarshal([]byte(recordString),&jsonMap)
	check(err)
	return jsonMap
}

func ReadEntryCsv() []ARequest {
	var reqS []ARequest
	dir :=GetRootFolder()
	entryCsv :=dir + "/entry.csv"
	file, err := os.Open(entryCsv)
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)

	for _,record:=range records{
		var req ARequest
		req.Url = record[0]
		req.ImgParam = record[1]
		req.ImgFolder = record[2]
		req.OutTime,_=strconv.Atoi(record[3])
		req.Headers = stringToMap(record[4])
		req.Params = stringToMap(record[5])

		reqS=append(reqS, req)
	}
	return reqS
}
