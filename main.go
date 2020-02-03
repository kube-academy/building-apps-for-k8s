package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Info *log.Logger
)

func initLog(infoHandle io.Writer) {
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	initLog(os.Stdout)

	Info.Println("Developer Productivity app starting")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	Info.Printf("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Developer Productivity app says Hi")
}
