package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Println("some one visted to the page")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>WELCOME  APP 1</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "to contact us mail @ <a href = \"akkutyagi.85.2@gmail.com\">Aakash</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Page not found</h1><p>Hi i am here come to me</p>")

	}

}

func handlerFuncInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO WEB APP INFO PAGE</h1>")
}

func main() {
	fmt.Println("some one visted to the page 0")
	http.HandleFunc("/", handlerFunc)
	fmt.Println("some one visted to the page 1")
	http.HandleFunc("/info", handlerFuncInfo)
	fmt.Println("some one visted to the page 2")
	http.ListenAndServe(":3000", nil)

}
