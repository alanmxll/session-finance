package http

import (
	"net/http"

	"github.com/alanmxll/session-finance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.ListenAndServe(":8080", nil)
}
