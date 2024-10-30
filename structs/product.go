package structs

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func Printprods(datas []Product) {
	for _, data := range datas {
		println("name :", data.Name, "Price:", data.Price, "Available", data.Available)
	}
}
