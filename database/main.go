package database

import (
	"database/sql"
	"fmt"

	"github.com/alanmxll/session-finance/model/transaction"
	"github.com/alanmxll/session-finance/util"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "financedb"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var db, _ = sql.Open("postgres", psqlInfo)
	return db
}

func Create(transaction transaction.Transaction) int {
	var db = connect()
	defer db.Close()

	var sqlStatement = `INSERT INTO transactions (title, amount, type, installment, created_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	// TODO: transaction.Installment
	var id int
	_ = db.QueryRow(sqlStatement,
		transaction.Title,
		transaction.Amount,
		transaction.Type,
		"",
		transaction.CreatedAt,
	)

	fmt.Println("New record ID is:", id)

	return id
}

func main() {
	// TODO: Installment
	Create(transaction.Transaction{
		Title:     "Another Freela",
		Amount:    600.0,
		Type:      0,
		CreatedAt: util.StringToTime("2022-01-03T21:05:00"),
	})

	fmt.Print(FetchAll())
}

func FetchAll() transaction.Transactions {
	var db = connect()
	defer db.Close()

	rows, _ := db.Query("SELECT title, amount, type, installment, created_at FROM transactions;")
	defer rows.Close()

	var transactionSlice []transaction.Transaction
	for rows.Next() {
		var transaction transaction.Transaction
		_ = rows.Scan(&transaction.Title,
			&transaction.Amount,
			&transaction.Type,
			"",
			&transaction.CreatedAt)

		transactionSlice = append(transactionSlice, transaction)
	}

	return transactionSlice
}
