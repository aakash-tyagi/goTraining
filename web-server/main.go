package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Println("some one visted to the page")
	fmt.Fprint(w, "<h1>WELCOME  APP</h1>")
}

func handlerFuncInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO WEB APP INFO PAGE</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/info", handlerFuncInfo)
	http.ListenAndServe(":3000", nil)
}
