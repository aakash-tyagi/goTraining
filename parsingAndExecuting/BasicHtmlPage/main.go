package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("BasicHtmlPage/*")) //--------used to parse file at folder level
}

func main() {
	// name := "aakash tyagi"

	tpl, err := tpl.ParseFiles("aaku.html") //----was used to parse files by name
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "aaku.html", nil) //-----execute specific template
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil) //------execute the first template it gets
	if err != nil {
		log.Fatal(err)
	}

}
