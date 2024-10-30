package main

import (
	_ "github.com/lib/pq"
	"github.com/sumit-behera-in/go-db-code/postgres"
)

func main() {

	db, err := postgres.Dbinitalizer()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	println("Creating a table if not exist")
	err = postgres.CreateProductTable(db, "table2")

	if err != nil {
		panic(err)
	}

}
