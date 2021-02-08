package main

import (
	"github.com/Golang-utils/websockets"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

//catch the signal ctrl+c

func catchSig(ch chan os.Signal, c *websocket.Conn) {
	//block in waiting for signal
	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "closed from client 1"))
	if err != nil {
		logrus.Error(err)
	}
	return
}

func main() {
	//connect the os signal to our channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	//use ws:// scheme to connect to websocket
	u := "ws://localhost:8000"
	logrus.Info("connecting with server...")
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	defer c.Close()
	//call catchSig in a goroutine
	go catchSig(interrupt, c)
	//start process
	websockets.Process(c)
}
