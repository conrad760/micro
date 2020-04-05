package main

import (
	"log"
	"net/http"
	"os"

	"github.com/conrad760/micro/handlers"
)

func main() {
	l := log.New(os.Stdout, "conrad-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
