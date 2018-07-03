package imgUpload

import (
	"os"
	"time"
)

func WriteToFile(text string) {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	_, err = f.WriteString(text+"\n")
	check(err)
	return
}

func WriteTime() {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	text := tm.Format("2006-01-02 03:04:05 PM")
	WriteToFile(text)
	return
}