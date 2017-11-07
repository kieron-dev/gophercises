package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestQuiz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quiz Suite")
}

var pathToQuiz string

var _ = BeforeSuite(func() {
	var err error
	pathToQuiz, err = gexec.Build("github.com/kieron-pivotal/gophercises/quiz")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
