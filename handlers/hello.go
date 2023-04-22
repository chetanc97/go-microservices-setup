package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, " Oops", http.StatusBadRequest)
		// w.WriteHeader(http.StatusBadRequest) "Alternative"
		// w.Write([]byte("Oops ! "))
		return
	}
	log.Printf("Data %s \n", d)
	fmt.Fprintf(w, "Returned back to user %s", d)
}
