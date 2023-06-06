package model

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func GetAllProducts() []Product {
	products := []Product{
		{
			ID:          1,
			Name:        "Nintendo Switch",
			Description: " Neon Blue and Neon Red",
			Price:       289.0,
		},
		{
			ID:          2,
			Name:        "Wireless Earbuds Bluetooth",
			Description: "5.3 Headphones 60 Hrs Playtime Sports",
			Price:       36.99,
		},
	}

	return products
}
