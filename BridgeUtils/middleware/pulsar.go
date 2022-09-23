package middleware

import (
	"log"
	"math/big"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type TxnType int

const (
	Warden TxnType = iota
	Enclave
	WardenTimeout
)

type OnboardTxnJSON struct {
	Type            TxnType  `json:"type"`
	BlockHash       string   `json:"blockHash"`
	TxnHash         string   `json:"txnHash"`
	ContractAddress string   `json:"contractAddress"`
	AccountAddress  string   `json:"accountAddress"`
	ChainId         *big.Int `json:"chainId"`
	BlockNumber     uint64   `json:"blockNumber"`
	TxnIndex        uint     `json:"txnIndex"`
	Amount          *big.Int `json:"amount"`
	OnboardTxnHash  string   `json:"onboardTxnHash"`
	Nonce           uint64   `json:"nonce"`
	Batch           int64    `json:"batch"`
}

type SettingsField struct {
	Path  []string `json:"path"`
	Value string   `json:"value"`
	Type  string   `json:"type"`
}

type SettingsJSON struct {
	Fields []*SettingsField `json:"fields"`
}

var (
	OnboardTxnSchemaDef = "{\"type\":\"record\",\"name\":\"OnboardTxn\"," +
		"\"fields\":[{\"name\":\"BlockHash\",\"type\":\"string\"}," +
		"{\"name\":\"Type\",\"type\":\"int\"}," +
		"{\"name\":\"TxnHash\",\"type\":\"string\"}," +
		"{\"name\":\"ContractAddress\",\"type\":\"string\"}," +
		"{\"name\":\"AccountAddress\",\"type\":\"string\"}," +
		"{\"name\":\"ChainId\",\"type\":\"int\"}," +
		"{\"name\":\"BlockNumber\",\"type\":\"int\"}," +
		"{\"name\":\"TxnIndex\",\"type\":\"int\"}," +
		"{\"name\":\"Amount\",\"type\":\"int\"}," +
		"{\"name\":\"OnboardTxnHash\",\"type\":\"string\"}," +
		"{\"name\":\"Nonce\",\"type\":\"int\"}," +
		"{\"name\":\"Batch\",\"type\":\"int\"}" +
		"]}"

	SettingsSchemaDef = `{"type":"record","name":"Settings",
		"fields":[{"name":"Fields","type":{"type":"array",
		"items":{"name":"Field","type":"record","fields":[{"name":"Value","type":"string"},
		{"name":"Type","type":"string"},
		{"name":"Path","type":{"type":"array","items":"string"}}]}}}]}`
)

func CreatePulsarClient() pulsar.Client {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               os.Getenv("PulsarURL"),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instance Pulsar client: %v", err)
	}
	return client
}
