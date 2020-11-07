package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"webalert/feed"
	"webalert/send"
)

var templates = template.Must(template.ParseFiles("tmpl/feed.html", "tmpl/send.html"))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/feed", http.StatusFound)
}

func feedHandler(w http.ResponseWriter, _ *http.Request) {
	err := templates.ExecuteTemplate(w, "feed.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sendHandler(w http.ResponseWriter, _ *http.Request) {
	err := templates.ExecuteTemplate(w, "send.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/feed", feedHandler)
	http.Handle("/events", broker)
	http.HandleFunc("/send", sendHandler)
	http.Handle("/socket", receiver)

	go broker.ServeMessages()
	go cmdMessenger(messages)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
