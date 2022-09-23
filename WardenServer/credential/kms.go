package credential

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	log "github.com/sirupsen/logrus"
)

type KmsCredentialFactory struct {
}

func (kcf *KmsCredentialFactory) MakeCredential() CredentialInterface {
	return &KmsCredential{
		AWSAccessKeyId:     os.Getenv("AWSAccessKeyId"),
		AWSAccessKeySecret: os.Getenv("AWSAccessKeySecret"),
		AWSRegion:          os.Getenv("AWSRegion"),
		AWSKMSKeyId:        os.Getenv("AWSKMSKeyId"),
	}
}

type KmsCredential struct {
	AWSAccessKeyId     string
	AWSAccessKeySecret string
	AWSRegion          string
	AWSKMSKeyId        string
}

type kmsDecryptAPI interface {
	Decrypt(ctx context.Context,
		params *kms.DecryptInput,
		optFns ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

func decryptData(c context.Context, api kmsDecryptAPI, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	return api.Decrypt(c, input)
}

func (kc *KmsCredential) Decrypt(ciphertext string) string {
	// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	// https://aws.github.io/aws-sdk-go-v2/docs/code-examples/kms/encryptdata/
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(kc.AWSAccessKeyId, kc.AWSAccessKeySecret, "")),
		config.WithRegion(kc.AWSRegion),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := kms.NewFromConfig(cfg)
	blob, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	input := &kms.DecryptInput{
		CiphertextBlob: blob,
	}

	decryptResult, err := decryptData(context.TODO(), client, input)
	if err != nil {
		log.Fatalf("kms decrypt err: %v", err)
	}
	return string(decryptResult.Plaintext)
}

type kmsEncryptAPI interface {
	Encrypt(ctx context.Context,
		params *kms.EncryptInput,
		optFns ...func(*kms.Options)) (*kms.EncryptOutput, error)
}

func encryptText(c context.Context, api kmsEncryptAPI, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
	return api.Encrypt(c, input)
}

func (kc *KmsCredential) Encrypt(plaintext string) string {
	// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	// https://aws.github.io/aws-sdk-go-v2/docs/code-examples/kms/encryptdata/

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(kc.AWSAccessKeyId, kc.AWSAccessKeySecret, "")),
		config.WithRegion(kc.AWSRegion),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := kms.NewFromConfig(cfg)

	input := &kms.EncryptInput{
		KeyId:     &kc.AWSKMSKeyId,
		Plaintext: []byte(plaintext),
	}

	result, err := encryptText(context.TODO(), client, input)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(result.CiphertextBlob)
}
