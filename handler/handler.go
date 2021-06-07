package handler

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Tree struct {
	Name string
}

func TemplateHandler(tmpl *template.Template) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		tree := Tree{
			Name: req.URL.Query().Get("favoriteTree"),
		}

		tmpl.Execute(w, tree)
	}
}

func CreateTemplate() *template.Template {
	content, err := ioutil.ReadFile("template.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("test").Parse(string(content))

	if err != nil {
		log.Fatal(err)
	}

	return tmpl
}
