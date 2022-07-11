package history

import (
	"encoding/json"
	"log"
	"net/http"

	"../account"
	"../web3/operations/subscribe"
)

// HistoryHendler ...
func HistoryHendler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientAddress := r.Form.Get("address")
	_, ok := account.Addresses[clientAddress]
	if ok {
		if _, ok := subscribe.AddressTransactions[clientAddress]; ok {
			jsonStr, err := json.Marshal(subscribe.AddressTransactions[clientAddress])
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write(jsonStr)
			return
		}
		w.Write([]byte("Not Data"))
		return
	}
	w.Write([]byte("Account not subsribe"))
	return

}
