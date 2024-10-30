package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/sumit-behera-in/go-db-code/postgres"
	"github.com/sumit-behera-in/go-db-code/structs"
)

func main() {
	postgresmain()
}

func postgresmain() {

	db, err := postgres.Dbinitalizer()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	println("Creating a table if not exist")
	err = db.CreateProductTable("dtable")
	if err != nil {
		log.Println(err)
	}

	println("Inserting it to db")
	var pk int
	pk, err = db.Insert("dtable", structs.Product{
		Name:      "P001",
		Price:     2433.09,
		Available: true,
	})
	if err != nil {
		log.Println(err)
	}
	println("The product is now inserted with id", pk)

	var dd []structs.Product
	dd, err = db.GetALL("dtable")
	if err != nil {
		log.Println(err)
	}
	structs.Printprods(dd)

	println("Update testig on id 2")
	var re int64
	prod := structs.Product{
		Name:      "ho",
		Available: false,
		Price:     66,
	}
	re, err = db.UpdateBYID(5, "dtable", prod)
	if err != nil {
		log.Println(err)
	}
	println(re, "lines are effected")

	println("Geting data with id 2")
	var data structs.Product
	data, err = db.GetRowByID(3, "dtable")
	if err != nil {
		log.Println(err)
	}
	println("name :", data.Name, "Price:", data.Price, "Available", data.Available)

	println("Deleting")
	re, err = db.DeleteById(2, "dtable")
	if err != nil {
		log.Println(err)
	}
	println(re, "lines are effected")

	println("Deleting")
	re, err = db.DeleteObject("dtable", structs.Product{
		Name:      "P001",
		Price:     2433.09,
		Available: true,
	})
	if err != nil {
		log.Println(err)
	}
	println(re, "lines are effected")

}
