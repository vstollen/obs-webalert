package feed

import (
	"fmt"
	"log"
)

type Broker struct {
	// Channel into which messages are pushed to be broadcast out to attached clients.
	Messages chan string

	// Set of clients, that should be messaged with incoming Messages.
	// The Set is represented by a map of empty structs. To add entries do:
	//   clients[chan] = struct{}{}
	clients map[client]struct{}
}

type client chan string

// ServeMessages waits for Messages and distributes these between all clients.
// Ends when Messages is closed.
func (b *Broker) ServeMessages() {
	for {
		msg, open := <-b.Messages

		if !open {
			log.Printf("The Messages channel was closed. Stopped message serving.")
			break
		}

		for client := range b.clients {
			client <- msg
		}
		log.Printf("Broadcast message to %d clients", len(b.clients))
	}
}

func (b *Broker) registerClient() client {

	if b.clients == nil {
		b.clients = make(map[client]struct{})
	}

	client := make(client, 1)

	b.clients[client] = struct{}{}

	fmt.Printf("New client registered.\n")

	return client
}

func (b *Broker) removeClient(client client) {
	delete(b.clients, client)
	close(client)

	fmt.Printf("Removed one client.\n")
}
