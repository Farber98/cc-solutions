package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8081", "specify port to listen on")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server on port %v!\n", *port)
	})

	log.Printf("Server listening on %v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
