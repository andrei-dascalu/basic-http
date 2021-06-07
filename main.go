package main

import (
	"log"
	"net/http"

	"github.com/andrei-dascalu/basic-http/handler"
)

func main() {
	http.HandleFunc("/favoriteTree", handler.TemplateHandler(handler.CreateTemplate()))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
