package main

import (
	"flag"
	"fmt"

	"rsc.io/quote"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" { // [O1o1o0g3o0]
		fmt.Println("Hello, World!")

		//} else if name == "hello" { // [O1o1o0g5o2o0]
		//g.Hello("Nanashino Gonbee")

		// この上に分岐を挟んでいく

	} else if name == "quote" { // [O1o1o0g4o0]
		fmt.Println(quote.Go())

	} else {
		fmt.Println("go run . {programName}")
	}
}
