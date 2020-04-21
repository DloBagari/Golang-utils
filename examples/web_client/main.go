package main

import (
	"github.com/Golang-utils/web_client"
	"github.com/sirupsen/logrus"
)

func main() {
	cli := web_client.Setup(true, false)
	if err := web_client.UseDefaultClient(); err != nil {
		logrus.Fatal(err)
	}
	if err := web_client.DoOps(cli); err != nil {
		logrus.Fatal(err)
	}
	c := web_client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		logrus.Info(err)
	}

	// create client with sure and noop, this will modify the http default client
	web_client.Setup(true, true)
	if err := web_client.UseDefaultClient(); err != nil {
		logrus.Info(err)
	}
}
