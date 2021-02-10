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
	}

}

func handlerFuncInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO WEB APP INFO PAGE</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/info", handlerFuncInfo)
	http.ListenAndServe(":3000", nil)
}
