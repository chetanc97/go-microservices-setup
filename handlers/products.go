package handlers

import (
	"encoding/json"
	"handlers/hello/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshall json ", http.StatusInternalServerError)
	} else {
		rw.Write([]byte(d))
	}
}
