package feed

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
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
		// Remove the client when the connection is closed.
		// When the connection is closed by the client, conn.ReadMessage() returns an error.
		// Because we don't expect the message feed to send data, we should only be able to read
		// on close.
		_, _, err := conn.ReadMessage()
		if err == nil {
			log.Println("Received unexpected message by client. Closing connection.")
		}
		b.removeClient(client)
	}()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(_ string) error {
		_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case msg, open := <-client:

			if !open {
				break
			}

			err := conn.WriteMessage(msg.MessageType, msg.Message)
			if err != nil {
				log.Println(err)
			}
			case <- ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
		}
	}
}
