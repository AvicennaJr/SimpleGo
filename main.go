package main

import (
	"flag"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	host := flag.String("host", "0.0.0.0", "Host to listen on")
	port := flag.Int("port", 8000, "Port to listen on")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	http.HandleFunc("/", helloHandler)

	fmt.Printf("Server starting on %s...\n", addr)
	http.ListenAndServe(addr, nil)
}
