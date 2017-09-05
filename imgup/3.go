package main

import (
	"net/http"
	"bytes"
	"mime/multipart"
)

func imgupload(url,err error){
	var b bytes.Buffer
	w:=multipart.NewWriter(&b)

}