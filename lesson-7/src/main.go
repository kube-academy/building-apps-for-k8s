package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "1.1"

func main() {
	log.Printf("[v%s] Starting my app", version)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "[v%s] My app says Hi!!!\n", version)
}
