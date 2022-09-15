package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", createAscii)
	mux.HandleFunc("/export", exportFile)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("web server is on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
