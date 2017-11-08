package main_test

import (
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
})
