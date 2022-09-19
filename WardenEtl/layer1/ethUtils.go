package layer1

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Addr = common.Address

type BoundContract struct {
	*bind.BoundContract
	addr Addr
	abi  string
	conn *ethclient.Client
}

func NewBoundContract(
	conn *ethclient.Client,
	addr Addr,
	rawABI string) (*BoundContract, error) {
	parsedABI, err := abi.JSON(strings.NewReader(rawABI))
	return &BoundContract{
		bind.NewBoundContract(addr, parsedABI, conn, conn, conn),
		addr,
		rawABI,
		conn,
	}, err
}

func (c *BoundContract) ParseEvent(
	name string,
	log types.Log,
	event interface{}) error {
	err := c.UnpackLog(event, name, log)
	return err
}

func MakeFilterQuery(addrs []common.Address, rawABI string, eventName string, fromBlock *big.Int, toBlock *big.Int, otherTopics [][]common.Hash) (ethereum.FilterQuery, error) {
	var q ethereum.FilterQuery

	parsedABI, err := abi.JSON(strings.NewReader((rawABI)))
	if err != nil {
		return q, err
	}

	ev, exist := parsedABI.Events[eventName]
	if !exist {
		return q, fmt.Errorf("unknown event name: %s", eventName)
	}

	q.FromBlock = fromBlock
	q.ToBlock = toBlock
	q.Addresses = addrs
	q.Topics = [][]common.Hash{{ev.ID}}
	q.Topics = append(q.Topics, otherTopics...)
	return q, nil
}
