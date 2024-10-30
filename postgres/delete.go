package postgres

import (
	"database/sql"
	"fmt"

	"github.com/sumit-behera-in/go-db-code/structs"
)

func DeleteById(db *sql.DB, id int, tableName string) (int64, error) {
	query := fmt.Sprintf(
		`DELETE FROM %s
        WHERE id = $1;`,
		tableName,
	)

	result, err := db.Exec(query, id)
	if err != nil {
		return 0, nil
	}

	var rowsEffected int64
	rowsEffected, err = result.RowsAffected()

	return rowsEffected, err
}

func DeleteObject(db *sql.DB, tableName string, product structs.Product) (int64, error) {
	query := fmt.Sprintf(
		`DELETE FROM %s
        WHERE name = $1, price = $2, available = $3`,
		tableName,
	)

	result, err := db.Exec(query, product.Name, product.Price, product.Available)
	if err != nil {
		return 0, nil
	}

	var rowsEffected int64
	rowsEffected, err = result.RowsAffected()

	return rowsEffected, err
}
