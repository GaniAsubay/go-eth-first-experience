package account

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/ethereum/go-ethereum/common"

	"../web3/operations/balance"
)

// Addresses ...
var Addresses map[string]string

func init() {
	Addresses = make(map[string]string)
}

// AddAccountHandler is ...
func AddAccountHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientAddress := r.Form.Get("address")
	re := regexp.MustCompile("^0x[0-9a-f]{40}$")
	if re.MatchString(clientAddress) {
		err := addAddressToMap(clientAddress)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
}

// GetEthBalanceHandler ...
func GetEthBalanceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientAddress := common.HexToAddress(r.Form.Get("address"))
	balance, err := balance.GetEthBalance(clientAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(balance))
	return
}

// GetERC20BalanceHandler ...
func GetERC20BalanceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientAddress := common.HexToAddress(r.Form.Get("address"))
	contractAddress := common.HexToAddress(r.Form.Get("token"))
	balance, err := balance.GetErc20Balance(clientAddress, contractAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(balance))
	return
}

// AddAddressToMap is ...
func addAddressToMap(address string) error {
	if _, ok := Addresses[address]; ok {
		return errors.New("Account before added")
	}
	Addresses[address] = address
	return nil
}
