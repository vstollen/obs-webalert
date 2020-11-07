package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"webalert/feed"
	"webalert/send"
)

func cmdMessenger(m chan string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("An Error occurred while reading your input:")
			log.Println(err.Error())
		}

		m <- message
	}
}

func main() {

	messages := make(chan string)
	broker := &feed.Broker{
		Messages: messages,
	}

	receiver := &send.Receiver{
		MessageSink: messages,
	}

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.Handle("/events", broker)
	http.Handle("/socket", receiver)

	go broker.ServeMessages()
	go cmdMessenger(messages)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
