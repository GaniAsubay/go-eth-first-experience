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
	Addresses["0x327b69f7086aa9eb478d5a871cf84aa3491f3a50"] = "0x327b69f7086aa9eb478d5a871cf84aa3491f3a50"
	Addresses["0x85f4c35de107ba8fb7a67f4348ed028d52c19254"] = "0x85f4c35de107ba8fb7a67f4348ed028d52c19254"
	Addresses["0x4bcb303609f19e71ab82a3a3393c46bfea1e44fc"] = "0x4bcb303609f19e71ab82a3a3393c46bfea1e44fc"
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
