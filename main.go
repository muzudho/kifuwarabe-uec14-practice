package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" {
		fmt.Println("Hello, World!")
	} else {
		fmt.Println("go run . {programName}")
	}
}
