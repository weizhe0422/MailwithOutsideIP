package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
)

func main() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		log.Fatalf("failed to query web page: %v",err)
	}
	defer resp.Body.Close()

	strIP, _ := ioutil.ReadAll(resp.Body)
	log.Println("IP:",strIP)

	from := "FROM_MAIL"
	pass := "MAIL_PASSWORD"
	to := "TO_MAIL"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		string(strIP)

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("OKOKOK!!")

}
