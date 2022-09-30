package credential

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestMakeCredential(t *testing.T) {

	t.Run("exist file", func(t *testing.T) {
		file := uuid.New().String()
		os.Setenv("CredentialFilePath", file)
		key := GenerateKey()
		ioutil.WriteFile(file, []byte(key), 0644)

		localCredentialFactory := &LocalCredentialFactory{}

		localCredential := localCredentialFactory.MakeCredential()
		if localCredential.GetKey() != key {
			t.Fatal("fail")
		}
		os.Remove(file)
	})

	t.Run("new file", func(t *testing.T) {
		file := uuid.New().String()
		os.Setenv("CredentialFilePath", file)

		localCredentialFactory := &LocalCredentialFactory{}
		localCredentialFactory.MakeCredential()

		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			t.Fatal("fail")
		}
		os.Remove(file)
	})

}
