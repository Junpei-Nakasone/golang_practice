package main

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.Writefile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.Readfile(filename)
	if err != nil {
		return
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "test", Body: []byte("this is")}
	p1.save()
}
