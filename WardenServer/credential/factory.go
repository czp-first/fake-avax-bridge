package credential

import (
	"fmt"
	"os"
)

type CredentialFactory interface {
	MakeCredential() CredentialInterface
}

type CredentialInterface interface {
	GetKey() string
	Encrypt(plaintext string) string
	Decrypt(ciphertext string) string
}

func GetCredential() (CredentialFactory, error) {

	credentialType := os.Getenv("CredentialType")
	if credentialType == "local" {
		return &LocalCredentialFactory{}, nil
	}

	// TODO
	// if credentialType == "aws" {
	// 	return
	// }

	return nil, fmt.Errorf("unknown credential type: %s", credentialType)
}
