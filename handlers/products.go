package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"example.com/hello/tutorial2/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(len(g))
		if len(g) != 1 {

			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.l.Println("Got id:", id)
		p.updateProduct(id, rw, r)

	}

	// catch all (else)
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()

	// Using json marshal
	// d, err := json.Marshal(lp)

	//  Or Using NewEncoder way
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	// If using json marshal
	// rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
