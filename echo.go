package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func echo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	fmt.Fprintf(w, strings.Replace(strings.Replace(r.URL.Path, "/", "", 1), "/", " ", -1))
}
