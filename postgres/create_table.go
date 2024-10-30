package postgres

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func CreateProductTable(db *sql.DB, name string) error {

	isValidName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	if !isValidName(name) {
		return fmt.Errorf("invalid table name: %s", name)
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`, name)

	// dont return anything
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
