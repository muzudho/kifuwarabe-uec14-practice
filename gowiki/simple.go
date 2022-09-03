// BOF [O1o1o0g9o7o_1o0]

package main

import (
	"fmt"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// EOF [O1o1o0g9o7o_1o0]
