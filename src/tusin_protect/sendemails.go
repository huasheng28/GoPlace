package tusin_protect

import (
	"net/smtp"
	"strings"
	"fmt"
)

func SendEmail(sender, pw string, receiver []string) bool {
	auth:=smtp.PlainAuth("",sender,pw,"smtp.mxhichina.com")
	to:=receiver
	nickname:=""
	user:=sender
	subject:="email head"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "email body."
	msg:=[]byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.mxhichina.com:25", auth, user, to, msg)
	if err!=nil{
		fmt.Println(err)
	}
}
