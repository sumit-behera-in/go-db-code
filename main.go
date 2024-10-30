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
		Price:     2433.09,
		Available: true,
	})
	if err != nil {
		panic(err)
	}
	println("The product is now inserted with id", pk)
	

	var dd []structs.Product
	dd, err = postgres.GetAllRowByName(db, "P001", "dtable")
	if err != nil {
		panic(err)
	}
	structs.Printprods(dd)

	println("Update testig on id 2")
	var re int64
	prod := structs.Product{
		Name:      "ho",
		Available: false,
		Price:     66,
	}
	re, err = postgres.UpdateBYID(db, 2, "dtable", prod)
	if err != nil {
		panic(err)
	}
	println(re, "lines are effected")

	println("Geting data with id 2")
	var data structs.Product
	data, err = postgres.GetRowByID(db, 2, "dtable")
	if err != nil {
		panic(err)
	}
	println("name :", data.Name, "Price:", data.Price, "Available", data.Available)

}
