package main

import (
	"fmt"
	"net/http"
)

type aakash int

func (m aakash) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is Online")
}

func main() {
	var tyagi aakash

	http.ListenAndServe(":9000", tyagi)
}
