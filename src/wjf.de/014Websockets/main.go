package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type message struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
}

func socket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message

		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}

		// send a response
		m2 := message{"Thanks for the message!"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}

		log.Println("Received message:", m.Message)

	}
}

func main() {
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	http.Handle("/socket", websocket.Handler(socket))
	http.ListenAndServe(":8080", nil)
}
