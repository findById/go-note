---
template:       article
title:          Golang smtp
date:           2016-12-01 15:00:00 +0800
keywords:       smtp
description:    golang send mail
---

```golang
package main

import (
    "net/smtp"
    "strings"
	"log"
)

func main() {
    user := "user@example.com"
	password := "password"
	addr := "mail.example.com"

	message := new(Message)
	message.sendTo = []string{"recipient@example.net"}
	message.subject = "Test subject"
	message.body = "This is the email body."

	SendMail(user, password, addr, message)
}

type Message struct {
	sendTo		[]string

	subject		string
	body		string
}

func SendMail(user, password, addr string, message *Message) {
	hostname := strings.Split(addr, ":")[0];

	// Set up authentication information.
	auth := smtp.PlainAuth("", user, password, hostname)

	msg := []byte("To: " + strings.Join(message.sendTo, ";") + "\r\n" +
		"From: " + user + "\r\n" +
		"Subject: " + message.subject + "\r\n" +
		"\r\n" + message.body + "\r\n")
	err := smtp.SendMail(addr, auth, user, message.sendTo, msg)
	if err != nil {
		log.Fatal(err)
	}
}
```