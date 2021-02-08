package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
)

//upgrader takes an http request and convert it to a websocket one.
var upgrader = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	//upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("failed in upgrading the connection", err)
		return
	}
	// read all messages
	for {
		messageType, byt, err := conn.ReadMessage()
		if err != nil {
			if strings.Contains(string(err.Error()), "close 1006") {
				logrus.Info("client closed")
			} else {
				fmt.Println(err.Error())
			}
			return
		}
		log.Printf("Recevied from client %#v", string(byt))
		//write back
		if err := conn.WriteMessage(messageType, byt); err != nil {
			log.Println("error in sending data to client", err)
		}

	}
}
