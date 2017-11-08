package question

import (
	"encoding/csv"
	"io"
	"os"
)

func LoadQuestions(filePath string) ([]QAPair, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(file)

	questions := []QAPair{}
	for {
		line, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		questions = append(questions, QAPair{
			Question: line[0],
			Answer:   line[1],
		})
	}

	return questions, nil
}
