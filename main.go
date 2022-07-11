package main

import (
	"fmt"
	"net/http"

	"./middleware"
	"./services/account"
	"./services/history"
	"./services/web3/operations/subscribe"
)

func init() {
	go subscribe.SubscribeNewHead()
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/subscribe", middleware.CheckAccountValid(http.HandlerFunc(account.AddAccountHandler)))
	mux.Handle("/history", middleware.CheckAccountValid(http.HandlerFunc(history.HistoryHendler)))
	mux.Handle("/balances", middleware.CheckAccountValid(http.HandlerFunc(account.GetEthBalanceHandler)))
	mux.Handle("/balance", middleware.CheckAccountValid(http.HandlerFunc(account.GetERC20BalanceHandler)))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
