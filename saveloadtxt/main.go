package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := &Page{Title: "FirstPage", Body: []byte("Here is the first page!")}
	p1.save()
	p2 := loadPage("FirstPage")
	fmt.Println(string(p2.Body))
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return &Page{Title: title, Body: body}
}
