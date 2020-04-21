package main

import (
	"github.com/Golang-utils/websockets"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.Info(" listening on port: 8000")
	//mount the handler on port 8000 to handle all request
	if err := http.ListenAndServe("localhost:8000", http.HandlerFunc(websockets.WSHandler)); err != nil {
		logrus.Fatal(err)
	}
}
