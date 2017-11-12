package question

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func LoadQuestions(filePath string, shuffle bool) ([]QAPair, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Question file '%s' cannot be found", filePath)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	questions := []QAPair{}
	for {
		line, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("Question file '%s' is not valid CSV", filePath)
		}
		if len(line) != 2 {
			return nil, fmt.Errorf("Question file '%s' must contain two columns", filePath)
		}
		questions = append(questions, QAPair{
			Question: line[0],
			Answer:   line[1],
		})
	}
	if shuffle {
		questions = shuffleQuestions(questions)
	}
	return questions, nil
}

func shuffleQuestions(questions []QAPair) []QAPair {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := []QAPair{}
	for _, i := range r.Perm(len(questions)) {
		ret = append(ret, questions[i])
	}
	return ret
}
