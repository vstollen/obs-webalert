package send

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Receiver struct {
	MessageSink chan string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (r *Receiver) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if messageType != websocket.TextMessage {
			errorMessage := fmt.Sprintf("Message Type: %v not supported!", messageType)
			panic(errorMessage)
		}

		r.MessageSink <- string(message)
	}
}