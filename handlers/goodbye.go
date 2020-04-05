package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// Goodbye introduce a logger for logging
type Goodbye struct {
	l *log.Logger
}
// NewGoodbye creates a new goodbye struct with a logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	g.l.Println("It was really great seeing you!")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s", d)
}
