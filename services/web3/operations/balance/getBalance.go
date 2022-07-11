package balance

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"../../../web3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// GetEthBalance ...
func GetEthBalance(account common.Address) (string, error) {
	client := web3.GetClient()
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", err
	}
	result := fromWai(balance)
	return result, nil
}

// GetErc20Balance ...
func GetErc20Balance(account common.Address, contractAddress common.Address) (string, error) {
	client := web3.GetClient()
	instance, _ := web3.NewTokenCaller(contractAddress, client)
	balance, err := instance.BalanceOfCustom(&bind.CallOpts{}, account)
	if err != nil {
		return "", err
	}

	return balance, nil
}

func fromWai(value *big.Int) string {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(18)))
	num, _ := decimal.NewFromString(value.String())
	return num.Div(mul).String()
}
