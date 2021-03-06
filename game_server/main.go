package main

import (
	"log"
	"net/http"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	consumableSquares := GenerateConsumableSquares(100)

	hub := newHub(consumableSquares)
	go hub.run()

	log.Println("Server started on localhost:3000")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
