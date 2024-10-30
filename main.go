package main

import (
	_ "github.com/lib/pq"
	"github.com/sumit-behera-in/go-db-code/postgres"
	"github.com/sumit-behera-in/go-db-code/structs"
)

func main() {

	db, err := postgres.Dbinitalizer()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	println("Creating a table if not exist")
	err = postgres.CreateProductTable(db, "dtable")
	if err != nil {
		panic(err)
	}

	println("Inserting it to db")
	var pk int
	pk, err = postgres.Insert(db, "dtable", structs.Product{
		Name:      "P001",
		Price:     24.09,
		Available: true,
	})
	if err != nil {
		panic(err)
	}
	println("The product is now inserted with id", pk)

}
