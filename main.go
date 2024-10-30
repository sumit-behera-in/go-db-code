package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "demo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	err = dbinitalizer(db, dbName)
	if err != nil {
		panic(err)
	}

	db.Close()

	fmt.Println("Reconneting to the database :", dbName)

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = createProductTable(db, "demoTable")
	if err != nil {
		panic(err)
	}
}

func dbinitalizer(db *sql.DB, dbName string) error {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	if err := db.QueryRow(query).Scan(&exists); err != nil {
		return fmt.Errorf("error checking if database exists: %v", err)
	}

	if !exists {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if _, err := db.Exec(createQuery); err != nil {
			return fmt.Errorf("unable to create database: %v", err)
		}
		fmt.Printf("Database %s created successfully!\n", dbName)
	} else {
		fmt.Printf("Database %s already exists.\n", dbName)
	}

	return nil
}

func createProductTable(db *sql.DB, name string) error {

	isValidName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	if !isValidName(name) {
		return fmt.Errorf("invalid table name: %s", name)
	}

	query := `CREATE TABLE IF NOT EXISTS dtable (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	// dont return anything
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
