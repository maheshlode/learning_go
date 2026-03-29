package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/maheshlode/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello1")

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT Request", r.URL)
		p.l.Println("URL Path", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -2)

		if len(g) != 1 {
			p.l.Println("Invalid len of g")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid g[0]")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Unable to convert to number", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("Got ID:", id)

		p.updateProducts(id, rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	log.Println("Hello2")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	// d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		p.l.Println("Error decoding JSON:", err)
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("UPDATE Request")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		p.l.Println("Error decoding JSON:", err)
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		p.l.Println("id: ", id)
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	// SUCCESS
	rw.WriteHeader(http.StatusNoContent)

}
