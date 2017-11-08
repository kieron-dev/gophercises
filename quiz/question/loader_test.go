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
		Expect(questions[0]).To(Equal(question.QAPair{
			Question: "5+5",
			Answer:   "10",
		}))
	})
})
