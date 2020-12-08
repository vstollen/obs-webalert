package feed

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// ServeHTTP handles websocket connections for the message feed.
// It registers the new client with the Broker and sends the client
// his messages.
func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := b.registerClient()

	go func() {
		<- r.Context().Done()
		b.removeClient(client)
	}()

	for {
		msg, open := <-client

		if !open {
			break
		}

		err := conn.WriteMessage(msg.MessageType, msg.Message)
		if err != nil {
			log.Println(err)
		}
	}
}
