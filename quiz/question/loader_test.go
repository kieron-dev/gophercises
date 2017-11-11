package question_test

import (
	"github.com/kieron-pivotal/gophercises/quiz/question"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Loader", func() {
	It("creates a list of questions from a CSV file", func() {
		questions, err := question.LoadQuestions("./fixtures/questions.csv")
		Expect(err).NotTo(HaveOccurred())
		Expect(len(questions)).To(Equal(13))
		Expect(questions).To(ContainElement(question.QAPair{
			Question: "5+5",
			Answer:   "10",
		}))
	})

	It("raises an error if a line does not contain the same number of items as the one above", func() {
		_, err := question.LoadQuestions("./fixtures/baddata.csv")
		Expect(err.Error()).To(ContainSubstring("not valid CSV"))
	})

	It("raises an error if the csv file doesn't contain a list of two item lines", func() {
		_, err := question.LoadQuestions("./fixtures/wrongdata.csv")
		Expect(err.Error()).To(ContainSubstring("must contain two columns"))
	})

	It("shuffles the questions", func() {
		questions1, _ := question.LoadQuestions("./fixtures/questions.csv")
		questions2, _ := question.LoadQuestions("./fixtures/questions.csv")
		Expect(questions1).NotTo(Equal(questions2))
	})
})
