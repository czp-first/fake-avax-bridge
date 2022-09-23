package credential

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type LocalCredentialFactory struct {
}

func (lcf *LocalCredentialFactory) MakeCredential() CredentialInterface {
	key, err := ioutil.ReadFile(os.Getenv("CredentialFilePath"))
	if err != nil {
		log.Fatal(err)
	}
	return &LocalCredential{Key: string(key)}
}

type LocalCredential struct {
	Key string `json:"key"`
}

func GenerateKey() string {
	keyBytes := make([]byte, 32)
	rand.Read(keyBytes)
	return base64.URLEncoding.EncodeToString(keyBytes)
}

type EncryptResponse struct {
	Ciphertext string `json:"ciphertext"`
}

func (lc *LocalCredential) Encrypt(plaintext string) string {

	requestBody := map[string]string{"key": lc.Key, "plaintext": plaintext}
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	url := "http://127.0.0.1:8050/encrypt"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var encryptResponse EncryptResponse
	err = json.NewDecoder(resp.Body).Decode(&encryptResponse)
	if err != nil {
		log.Fatal(err)
	}
	return encryptResponse.Ciphertext
}

type DecryptResponse struct {
	Plaintext string `json:"plaintext"`
}

func (lc *LocalCredential) Decrypt(ciphertext string) string {

	requestBody := map[string]string{"key": lc.Key, "ciphertext": ciphertext}
	json_data, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	url := "http://127.0.0.1:8050/decrypt"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var decryptResponse DecryptResponse
	err = json.NewDecoder(resp.Body).Decode(&decryptResponse)
	if err != nil {
		log.Fatal(err)
	}
	return decryptResponse.Plaintext
}
