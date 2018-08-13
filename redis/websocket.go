package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"log"
)

var upgrade = websocket.Upgrader{
	// cross origin domain
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/echo", echo)
	http.ListenAndServe("localhost:8080", nil)
	log.Fatal("start")
}


func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	for {
		mt, message, _ := c.ReadMessage()
		fmt.Println(mt, string(message))
		c.WriteMessage(mt, append([]byte("hello "), message[:]...))
	}

}
