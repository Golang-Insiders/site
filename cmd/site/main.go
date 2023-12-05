package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("failed to serve home page.")
	}	
}
