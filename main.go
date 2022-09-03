package main

import (
	"flag"
	"fmt"

	"github.com/muzudho/kifuwarabe-uec14-practice/greetings"
	"rsc.io/quote"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" { // [O1o1o0g3o0]
		fmt.Println("Hello, World!")

	} else if name == "greetings" { // [O1o1o0g5o2o0]
		fmt.Println(greetings.GetMessage("Nanashino Gonbee"))

	} else if name == "quote" { // [O1o1o0g4o0]
		fmt.Println(quote.Go())

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}
