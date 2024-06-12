package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type routeServe struct{}

func (rs *routeServe) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	w.Write([]byte(fmt.Sprintf("Hello %s", id)))
}

func main() {
	port := flag.Int("port", 12345, "Port number to be used for HTTP")
	flag.Parse()
	color.Cyan("Application is running on port %d\n", *port)

	servingHandler := &routeServe{}
	httpMux := http.NewServeMux()
	httpMux.Handle("GET /hello/{id}", servingHandler)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: httpMux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error running HTTP: %v\n", err)
	}
}
