package chainlinkfeed

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func init() {
	os.Setenv("MainnetETHHttps", "https://mainnet.infura.io/v3/e82271b457844bca96ea810c5947e989")
}

func TestGetFeedData(t *testing.T) {
	// aave: 0x6df09e975c830ecae5bd4ed9d90f3a95a4f88012
	aaveContract := common.HexToAddress("0x6df09e975c830ecae5bd4ed9d90f3a95a4f88012")
	price, err := GetFeedData(aaveContract)
	fmt.Println(err)
	fmt.Println(price)
}
