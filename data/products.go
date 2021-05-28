package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

var ErrorProductNotFound = fmt.Errorf("Prodct Not Found")

func findProduct(id int) (int, error) {
	for idx, p := range productList {
		if p.ID == id {
			return idx, nil
		}
	}
	return -1, ErrorProductNotFound
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil
}

var loc, _ = time.LoadLocation("Asia/Kolkata")

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc242",
		CreatedOn:   time.Now().In(loc).String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and Strong coffee without milk",
		Price:       1.99,
		SKU:         "hdf921",
		CreatedOn:   time.Now().In(loc).String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
