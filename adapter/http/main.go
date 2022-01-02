package http

import (
	"net/http"

	"github.com/alanmxll/session-finance/adapter/http/actuator"
	"github.com/alanmxll/session-finance/adapter/http/transaction"
)

// Init function
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
