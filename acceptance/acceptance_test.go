package acceptance_test

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Acceptance", func() {
	var cmd *exec.Cmd

	BeforeEach(func() {
		cmd = exec.Command(binPath)
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		Expect(cmd.Start()).To(Succeed())
	})

	AfterEach(func() {
		Expect(cmd.Process.Signal(os.Interrupt)).To(Succeed())
		Expect(cmd.Wait()).To(Succeed())
	})

	Describe("GET /ping", func() {
		It("responds with 204 No Content", func() {
			Eventually(ping(serverAddress)).Should(Succeed())
		})
	})
})

func ping(url string) func() error {
	return func() error {
		response, err := http.Get(url + "/ping")
		if err != nil {
			return err
		}
		if err = response.Body.Close(); err != nil {
			return err
		}

		if response.StatusCode == http.StatusNoContent {
			return nil
		}

		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}
