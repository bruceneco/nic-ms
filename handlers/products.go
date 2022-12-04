package handlers

import (
	"github.com/bruceneco/nic-ms/data"
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
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w)
	}
}

func (p *Products) getProducts(w http.ResponseWriter) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
	}
}
