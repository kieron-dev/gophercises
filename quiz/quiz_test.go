package main_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Quiz", func() {

	Context("command flags", func() {
		It("displays usages given -h flag", func() {
			command := exec.Command(pathToQuiz, "-h")
			stderr := gbytes.NewBuffer()
			session, err := gexec.Start(command, GinkgoWriter, stderr)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(2))
			Eventually(stderr).Should(gbytes.Say("Usage"))
		})

		It("doesn't display usage when no -h flag", func() {
			command := exec.Command(pathToQuiz)
			stdout := gbytes.NewBuffer()
			session, err := gexec.Start(command, stdout, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit())
			Consistently(stdout).ShouldNot(gbytes.Say("Usage"))
		})
	})

	Context("csv file", func() {
		It("exits and says file does not exist if csv file cannot be found", func() {
			csvFile := "./fixtures/does-not-exist.csv"
			command := exec.Command(pathToQuiz, "-csv", csvFile)
			stderr := gbytes.NewBuffer()
			session, err := gexec.Start(command, GinkgoWriter, stderr)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Eventually(stderr).Should(gbytes.Say(fmt.Sprintf("Question file '%s' cannot be found", csvFile)))
		})

		It("exits and says file is invalid if columns are missing", func() {
			csvFile := "./question/fixtures/baddata.csv"
			command := exec.Command(pathToQuiz, "-csv", csvFile)
			stderr := gbytes.NewBuffer()
			session, err := gexec.Start(command, GinkgoWriter, stderr)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Eventually(stderr).Should(gbytes.Say(fmt.Sprintf("Question file '%s' is not valid CSV", csvFile)))
		})
	})
})
