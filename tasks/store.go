package tasks

import "fmt"

type Product struct {
	Name  string
	Price int
}

type Store struct {
	Products []Product
}

func (s *Store) AddProduct(name string, price int) {
	s.Products = append(s.Products, Product{
		Name:  name,
		Price: price,
	})
}

func (s *Store) FindByName(name string) (*Product, error) {
	for i := range s.Products {
		if s.Products[i].Name == name {
			return &s.Products[i], nil
		}
	}
	return nil, fmt.Errorf("product %q not found", name)
}

func (s *Store) TotalPrice() int {
	var total int
	for _, p := range s.Products {
		total += p.Price
	}
	return total
}
