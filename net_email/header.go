package net_email

import (
	"fmt"
	"net/mail"
	"strings"
)

func PrintHeaderInfo(header mail.Header) error {
	//this will work is email is sent to a single address
	// if email is been set to multiple addresses we need to use ParseAddressList
	fmt.Println(strings.Repeat("=", 40))
	toAddress, err := mail.ParseAddress(header.Get("To"))
	if err != nil {
		return err
	}
	fmt.Printf("email is sent to %s email address %s\n", toAddress.Name, toAddress.Address)
	fromAddress, err := mail.ParseAddress(header.Get("From"))
	if err != nil {
		return err
	}
	fmt.Printf("email is sent from %s address %s\n", fromAddress.Name, fromAddress.Address)
	fmt.Println("Subject:", header.Get("Subject"))

	//date will work if email has valid RFC5322 date
	if date, err := header.Date(); err == nil {
		fmt.Printf("Date: %s\n", date.String())
	}
	fmt.Println(strings.Repeat("=", 40))
	return nil
}
