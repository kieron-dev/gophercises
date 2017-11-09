package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kieron-pivotal/gophercises/quiz/question"
)

func main() {
	var (
		csv   string
		limit int
	)
	flag.StringVar(&csv, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&limit, "limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	questions, err := question.LoadQuestions(csv)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for i, q := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, q.Question)
		input := ""
		fmt.Scanln(&input)
	}
}
