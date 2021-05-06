package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, "yo.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

}
