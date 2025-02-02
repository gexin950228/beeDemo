package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type MailContext struct {
	From        string
	To          string
	Subject     string
	Body        string
	ContextType string
}

type SendResult struct {
	Code  int
	Error error
}

func SendMail(mailContext MailContext) SendResult {
	mailConfig := LoadConfig()
	msg := gomail.NewMessage()
	var sendResult SendResult
	msg.SetHeader("From", mailContext.From)
	msg.SetHeader("To", mailContext.To)
	msg.SetHeader("Subject", mailContext.Subject)
	msg.SetBody(mailContext.ContextType, mailContext.Body)
	dial := gomail.NewDialer(mailConfig.Host, mailConfig.Port, mailConfig.Account, mailConfig.Password)
	err := dial.DialAndSend(msg)
	if err != nil {
		LogToFile("Panic", fmt.Sprintf("发送邮件出错，错误信息: %s", err.Error()))
		sendResult = SendResult{Code: 0, Error: err}
	} else {
		sendResult = SendResult{Code: 1, Error: nil}
		LogToFile("Info", "没有错误")
	}
	return sendResult
}
