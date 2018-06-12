package tusin_protect

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"strconv"
)

func ReadCsv() (int, int, []string, []string, []string, []string, []string) {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	csvFile := currentPath + "/data.csv"
	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	i := 0
	var delaytimes, outtimes string
	var delaytime, outtime int
	var urlSlice, pathSlice, phoneSlice, emailSlice, receiverSlice []string
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			delaytimes = line[0]
			delaytime, _ = strconv.Atoi(delaytimes)
		} else if i == 1 {
			outtimes = line[0]
			outtime, _ = strconv.Atoi(outtimes)
		} else if i == 2 {
			urlSlice = line
		} else if i == 3 {
			pathSlice = line
		} else if i == 4 {
			phoneSlice = line
		} else if i == 5 {
			emailSlice = line
		} else if i == 6 {
			receiverSlice = line
		}
		i++
	}
	return delaytime, outtime, urlSlice, pathSlice, phoneSlice, emailSlice, receiverSlice
}
