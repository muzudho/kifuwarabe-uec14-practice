// BOF [O1o1o0g9o7o0]

package main

import (
	"fmt"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// EOF [O1o1o0g9o7o0]
