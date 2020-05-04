// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("listening on port 8080")
	err := http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./wasm`)))
	log.Fatalln(err)
}
