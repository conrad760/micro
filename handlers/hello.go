package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello introduces a new logger for logging
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello struct with a logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
