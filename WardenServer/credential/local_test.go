package credential_test

import (
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"

	"github.com/czp-first/fake-avax-bridge/WardenServer/credential"
)

var _ = Describe("Local credential", func() {

	Describe("MakeCredential", func() {
		var file, key string
		var localCredentialFactory *credential.LocalCredentialFactory

		BeforeEach(func() {
			file = uuid.New().String()
			key = credential.GenerateKey()
			os.Setenv("CredentialFilePath", file)
			localCredentialFactory = &credential.LocalCredentialFactory{}
		})

		AfterEach(func() {
			os.Remove(file)
		})

		Context("existed file", func() {
			It("read existed file", func() {
				ioutil.WriteFile(file, []byte(key), 0644)
				localCredential := localCredentialFactory.MakeCredential()
				Expect(key).To(Equal(localCredential.GetKey()))
			})
		})

		Context("new file", func() {
			It("create file", func() {
				localCredentialFactory.MakeCredential()
				_, err := os.Stat(file)
				Expect(os.IsNotExist(err)).To(BeFalse())
			})
		})
	})

	Describe("GetKey", func() {
		It("get key success", func() {
			key := "imkey"
			lc := &credential.LocalCredential{
				Key: key,
			}

			Expect(lc.GetKey()).To(Equal(key))
		})
	})

	Describe("Encrypt", func() {
		It("encrypt success", func() {
			defer gock.Off()

			cloudServiceUrl := "http://credential.com"
			os.Setenv("CloudServiceURL", cloudServiceUrl)

			gock.New(cloudServiceUrl).
				Post("/encrypt").
				Reply(200).
				JSON(map[string]string{"ciphertext": "encrypt-text"})
			lc := &credential.LocalCredential{
				Key: "key",
			}

			Expect(lc.Encrypt("text")).To(Equal("encrypt-text"))
		})
	})

	Describe("Decrypt", func() {
		It("decrypt", func() {
			defer gock.Off()

			cloudServiceUrl := "http://credential.com"
			os.Setenv("CloudServiceURL", cloudServiceUrl)

			gock.New(cloudServiceUrl).
				Post("/decrypt").
				Reply(200).
				JSON(map[string]string{"plaintext": "decrypt-text"})
			lc := &credential.LocalCredential{
				Key: "key",
			}

			Expect(lc.Decrypt("ciphertext")).To(Equal("decrypt-text"))
		})
	})
})
