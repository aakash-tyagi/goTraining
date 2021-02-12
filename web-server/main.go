package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFuncInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO WEB APP INFO PAGE</h1>")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO HOME</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>to contact us mail @ <a href = \"8010574809\">Aakash</a></h1>")
}

func mahno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>call me 8010574809</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/info", handlerFuncInfo)
	r.HandleFunc("/8010574809", mahno)

	http.ListenAndServe(":3000", r)

}
