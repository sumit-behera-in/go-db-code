package structs

type Product struct {
	Name      string  `bson:"name"`
	Price     float64 `bson:"price"`
	Available bool    `bson:"available"`
}

func Printprods(datas []Product) {
	for _, data := range datas {
		println("name :", data.Name, "Price:", data.Price, "Available", data.Available)
	}
}
