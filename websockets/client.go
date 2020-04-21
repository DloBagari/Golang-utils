package websockets

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Process(c *websocket.Conn) {
	//read from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your text: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		logrus.Fatal(err)
	}
	data = strings.TrimSpace(data)
	//send data to server
	if err := c.WriteMessage(websocket.TextMessage, []byte(data)); err != nil {
		logrus.Fatal(err)
	}
	//read data from server
	_, byt, err := c.ReadMessage()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("data from server : %#v", string(byt))
}
