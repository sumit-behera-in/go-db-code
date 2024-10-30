package postgres

import (
	"database/sql"
	"fmt"

	"github.com/sumit-behera-in/go-db-code/structs"
)

func Insert(db *sql.DB, tableName string, products structs.Product) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (name,price,available)
	VALUES ($1,$2,$3) RETURNING id`, tableName)

	// retuned id is pk
	var pk int
	err := db.QueryRow(query, products.Name, products.Price, products.Available).Scan(&pk)
	return pk, err
}
