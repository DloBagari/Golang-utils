package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	p := &Proxy{
		Client:  http.DefaultClient,
		BaseUrl: "https://www.golang.org",
	}
	http.Handle("/", p)
	logrus.Info("starting server ..")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		logrus.Fatal(err)
	}
}
