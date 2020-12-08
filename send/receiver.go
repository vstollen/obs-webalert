package send

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Message struct {
	MessageType int
	Message []byte
}

type Receiver struct {
	MessageSink chan Message
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

	fmt.Printf("New Sender registered.\n")

	for {
		messageType, messageData, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			MessageType: messageType,
			Message:     messageData,
		}

		r.MessageSink <- message
	}
}