package main

import (
	"bytes"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "t0031692@gmail.com", "satk afpv gxpn ldjc")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	email, err := d.Dial()
	if err != (nil) {
		panic(err)
	}
	for {
		var msg *bytes.Buffer = bytes.NewBuffer([]byte("GET FLOODED!!!!!!!!"))
		email.Send("t0031692@gmail.com", []string{"t0031692@gmail.com"}, msg)

	}
}
