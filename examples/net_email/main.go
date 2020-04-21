package main

import (
	"fmt"
	"github.com/Golang-utils/net_email"
	"io/ioutil"
	"net/mail"
	"strings"
)

func main() {
	s := `Date: Thu, 24 Jul 2020 08:00:32 -0700
From: sender <dlo@live.ie>
To: receiver <dlo@outlook.ie>
Subject: this is a subject title

this is a subject data
`
	m, err := mail.ReadMessage(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	err = net_email.PrintHeaderInfo(m.Header)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(m.Body)
	fmt.Println(string(data))

}
