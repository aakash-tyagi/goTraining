package main

import (
	"io"
	"log"
	"net/http"
)

func slashHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func dogHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func namePrinter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func main() {

	http.HandleFunc("/", slashHandle)
	http.HandleFunc("/dog", dogHandle)
	http.HandleFunc("/me/:", namePrinter)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
