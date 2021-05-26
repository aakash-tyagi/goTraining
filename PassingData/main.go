package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, "al jazera") //------execute the first template it gets
	if err != nil {
		log.Fatal(err)
	}
}
