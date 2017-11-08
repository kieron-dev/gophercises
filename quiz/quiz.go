package main

import (
	"flag"
)

func main() {
	var (
		csv   string
		limit int
	)
	flag.StringVar(&csv, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&limit, "limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
}
