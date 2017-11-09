package main_test

import (
	"fmt"
	"io"
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

	Context("asking questions", func() {
		var (
			command *exec.Cmd
			stdin   io.WriteCloser
			stdout  *gbytes.Buffer
			stderr  *gbytes.Buffer
			err     error
		)

		BeforeEach(func() {
			command = exec.Command(pathToQuiz, "-csv", "./question/fixtures/questions.csv")
			stdin, err = command.StdinPipe()
			stdout = gbytes.NewBuffer()
			stderr = gbytes.NewBuffer()
			_, err := gexec.Start(command, stdout, stderr)
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			gexec.Kill()
		})

		It("successfully asks the first question", func() {
			Eventually(stdout).Should(gbytes.Say(`Problem #1: `))
		})

		It("pauses after the first question waiting for input", func() {
			Consistently(stdout).ShouldNot(gbytes.Say(`Problem #2: `))
		})

		It("prints the second question after first answer entered", func() {
			fmt.Fprintln(stdin, "foo")
			Eventually(stdout).Should(gbytes.Say(`Problem #2: `))
		})
	})

	Context("keeping score", func() {
		var (
			command *exec.Cmd
			stdin   io.WriteCloser
			stdout  io.Writer
			stderr  io.Writer
			err     error
		)

		BeforeEach(func() {
			command = exec.Command(pathToQuiz, "-csv", "./question/fixtures/onequestion.csv")
			stdin, err = command.StdinPipe()
			stdout = gbytes.NewBuffer()
			stderr = GinkgoWriter
			_, err := gexec.Start(command, stdout, stderr)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			gexec.Kill()
		})

		It("should return 1 / 1 for a correct answer in 1 question test", func() {
			fmt.Fprintln(stdin, "10")
			Eventually(stdout).Should(gbytes.Say(`You scored 1 out of 1.`))
		})

		It("should return 0 / 1 for an incorrect answer in 1 question test", func() {
			fmt.Fprintln(stdin, "100")
			Eventually(stdout).Should(gbytes.Say(`You scored 0 out of 1.`))
		})
	})
})
