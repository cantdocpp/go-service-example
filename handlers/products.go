package handlers

import (
	// "encoding/json"
	"github.com/cantdocpp/go-service-example/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to marhsal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
