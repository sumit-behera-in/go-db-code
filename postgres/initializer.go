package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func Dbinitalizer() (DB, error) {

	db_instance := DB{}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db_instance, err
	}

	err = db.Ping()
	if err != nil {
		return db_instance, err
	}

	fmt.Println("Successfully connected!")

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	if err := db.QueryRow(query).Scan(&exists); err != nil {
		return db_instance, fmt.Errorf("error checking if database exists: %v", err)
	}

	if !exists {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if _, err := db.Exec(createQuery); err != nil {
			return db_instance, fmt.Errorf("unable to create database: %v", err)
		}
		fmt.Printf("Database %s created successfully!\n", dbName)
	} else {
		fmt.Printf("Database %s already exists.\n", dbName)
	}

	db.Close()

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return db_instance, err
	}

	db_instance.db = db

	return db_instance, nil
}

func (db *DB) Close() {
	db.db.Close()
}
