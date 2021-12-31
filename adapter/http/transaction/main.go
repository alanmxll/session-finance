package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/alanmxll/session-finance/model/transaction"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2021-12-26T18:24:16")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    2625.00,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
