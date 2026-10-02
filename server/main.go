package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Distributed Multiplayer Game Server is running!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	port := ":8080"

	log.Println("Game server starting...")
	log.Println("Server listening on http://localhost:8080")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}