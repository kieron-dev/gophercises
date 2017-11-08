package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "show usage")
	flag.Parse()

	usage := `Usage of %s:
  -csv string
        a csv file in the format of 'question,answer' (default "problems.csv")
  -limit int
        the time limit for the quiz in seconds (default 30)
`

	if help {
		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}
}
