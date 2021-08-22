package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	serverMux := http.NewServeMux()

	//web page
	serverMux.Handle("/", http.FileServer(http.Dir("public")))

	// server websocket
	serverMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(hub, w, r)
	})

	log.Println("running on port :8080")
	log.Println(http.ListenAndServe(":8080", serverMux))
}
