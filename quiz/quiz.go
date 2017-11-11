package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kieron-pivotal/gophercises/quiz/question"
)

func main() {
	var (
		csv   string
		limit float64
	)
	flag.StringVar(&csv, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Float64Var(&limit, "limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	questions, err := question.LoadQuestions(csv)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	score := 0
	timeout := make(chan interface{})

	go func() {
		time.Sleep(time.Duration(limit) * time.Second)
		fmt.Println("\nTime's up!")
		close(timeout)
	}()

	go func() {
		for i, q := range questions {
			fmt.Printf("Problem #%d: %s = ", i+1, q.Question)
			input := ""
			fmt.Scanln(&input)
			if input == q.Answer {
				score++
			}
		}
		close(timeout)
	}()

	<-timeout
	fmt.Printf("You scored %d out of %d.\n", score, len(questions))
}
