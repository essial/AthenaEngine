package aehttp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Serve() {
	fs := http.FileServer(http.Dir("./aehttp/www"))
	http.HandleFunc("/wsconnect", handleWebSocket)
	http.Handle("/", fs)
	log.Print("Http server listening on 3000")
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			//TODO: notify server
			log.Fatal(err)
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			//TODO: notify server
			return
		}
	}
}
