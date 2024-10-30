package postgres

import (
	"fmt"

	"github.com/sumit-behera-in/go-db-code/structs"
)

func (db *DB) GetRowByID(id int, tableName string) (structs.Product, error) {
	query := fmt.Sprintf(
		`SELECT name, price, available FROM %s WHERE id = $1`,
		tableName,
	)

	var name string
	var available bool
	var price float64

	err := db.db.QueryRow(query, id).Scan(&name, &price, &available)

	return structs.Product{
		Name:      name,
		Available: available,
		Price:     price,
	}, err
}

func (db *DB) GetAllRowByName(name_filed string, tableName string) ([]structs.Product, error) {
	query := fmt.Sprintf(
		`SELECT name, price, available FROM %s WHERE name = $1`,
		tableName,
	)

	var data []structs.Product
	rows, err := db.db.Query(query, name_filed)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var name string
	var available bool
	var price float64

	for rows.Next() {
		if err = rows.Scan(&name, &price, &available); err != nil {
			return data, err
		}
		data = append(data, structs.Product{
			Name:      name,
			Available: available,
			Price:     price,
		})
	}

	return data, err
}

func (db *DB) GetALL(tableName string) ([]structs.Product, error) {
	query := fmt.Sprintf(
		`SELECT name, price, available FROM %s `,
		tableName,
	)

	var data []structs.Product
	rows, err := db.db.Query(query)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var name string
	var available bool
	var price float64

	for rows.Next() {
		if err = rows.Scan(&name, &price, &available); err != nil {
			return data, err
		}
		data = append(data, structs.Product{
			Name:      name,
			Available: available,
			Price:     price,
		})
	}

	return data, err
}
