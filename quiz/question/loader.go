package question

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LoadQuestions(filePath string) ([]QAPair, error) {
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

	return questions, nil
}
