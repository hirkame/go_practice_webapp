package main

import (
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	// Import CSS
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Handler
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
