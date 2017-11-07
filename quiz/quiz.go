package main

import (
	"flag"
	"fmt"
)

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "show usage")
	flag.Parse()

	if help {
		fmt.Println("Usage: foo")
	}
}
