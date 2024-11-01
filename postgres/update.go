package postgres

import (
	"fmt"

	"github.com/sumit-behera-in/go-db-code/structs"
)

func (db *DB) UpdateBYID(id int, tableName string, product structs.Product) (int64, error) {
	query := fmt.Sprintf(
		`UPDATE %s
        SET name = $1, price = $2, available = $3
        WHERE id = $4;`,
		tableName,
	)

	result, err := db.db.Exec(query, product.Name, product.Price, product.Available, id)
	if err != nil {
		return 0, nil
	}

	var rowsEffected int64
	rowsEffected, err = result.RowsAffected()

	return rowsEffected, err
}
