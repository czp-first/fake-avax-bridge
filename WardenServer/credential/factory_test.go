package credential_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/czp-first/fake-avax-bridge/WardenServer/credential"
)

var _ = Describe("Factory", func() {
	Describe("GetCredentialFactory", func() {
		It("local", func() {
			os.Setenv("CredentialType", "local")
			credentialFactory, err := credential.GetCredentialFactory()
			Expect(credentialFactory).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("unknown type", func() {
			os.Setenv("CredentialType", "oss")
			credentialFactory, err := credential.GetCredentialFactory()
			Expect(credentialFactory).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
