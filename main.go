package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	log.Println("Started listening on port :8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
