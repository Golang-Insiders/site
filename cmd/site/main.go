package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("failed to serve home page.")
	}
}
