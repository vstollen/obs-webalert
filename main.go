package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"webalert/feed"
	"webalert/send"
)

func main() {

	messages := make(chan send.Message)
	broker := &feed.Broker{
		Messages: messages,
	}

	receiver := &send.Receiver{
		MessageSink: messages,
	}

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.Handle("/feed", broker)
	http.Handle("/socket", receiver)

	go broker.ServeMessages()
	go cmdMessenger(messages)

	fmt.Printf("Starting Server on Post 8080.\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// cmdMessenger forwards all messages from Stdin into the m channel.
func cmdMessenger(m chan send.Message) {
	reader := bufio.NewReader(os.Stdin)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("An Error occurred while reading your input:")
			log.Println(err.Error())
		}

		m <- send.Message{
			MessageType: websocket.TextMessage,
			Message:     []byte(message),
		}
	}
}
