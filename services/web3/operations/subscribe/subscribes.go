package subscribe

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"../../../account"
	"../../../web3"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Transaction ...
type Transaction struct {
	Sum      int64
	Txn      string
	IsSender bool
}

var (
	// AddressTransactions ...
	AddressTransactions map[string][]Transaction
	client              *ethclient.Client
)

func init() {
	AddressTransactions = make(map[string][]Transaction)
	client = web3.GetClient()
}

// SubscribeNewHead ...
func SubscribeNewHead() {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("SubscribeNewHead Err:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("Subs chan Err: ", err)
		case header := <-headers:
			time.Sleep(5 * time.Second)
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Println("Get Block By Hash Err: ", err, "Block Hash: ", header.Hash())
			} else {
				for _, tx := range block.Transactions() {
					txAsMessage, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(0))
					if err != nil {
						continue
					}
					to := txAsMessage.To()
					if to == nil {
						continue
					}
					toAddress := strings.ToLower(to.String())
					fromAddress := strings.ToLower(txAsMessage.From().String())
					if _, ok := account.Addresses[toAddress]; ok {
						data := Transaction{Sum: tx.Value().Int64(), Txn: tx.Hash().String(), IsSender: false}
						addTransaction(toAddress, data)
					}

					if _, ok := account.Addresses[fromAddress]; ok {
						data := Transaction{Sum: tx.Value().Int64(), Txn: tx.Hash().String(), IsSender: true}
						addTransaction(fromAddress, data)
					}
				}
			}
		}
	}
}

// AddTransaction ...
func addTransaction(address string, data Transaction) {
	AddressTransactions[address] = append(AddressTransactions[address], data)
}
