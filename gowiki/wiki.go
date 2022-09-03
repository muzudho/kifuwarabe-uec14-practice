// BOF [O1o1o0g9o5o0]

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "vol1" { // [O1o1o0g9o5o0]
		p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
		p1.save()
		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.Body))

	} else if name == "vol2" { // [O1o1o0g9o7o_2o0]
		http.HandleFunc("/", SimpleHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))

	} else if name == "vol3" { // [O1o1o0g9o8o0]
		http.HandleFunc("/view/", ViewHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o5o0]
