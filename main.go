package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
	host = "http://localhost"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: blue'>Hello World</h1>"))

}

func main() {
	http.HandleFunc("/hello", Hello)

	// log.Println("Server started at", "\033[1;32m"+host+port+"/hello", "\033[0m")
	log.Println("Server started at", host+port+"/hello")

	http.ListenAndServe(port, nil)
}
