package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := pat.New()

	mux.Get("/hello", http.HandlerFunc(hello))
	mux.Add("GET", "/", http.FileServer(http.Dir("./public")))

	http.Handle("/", mux)

	port := os.Getenv("PORT")

	log.Println("Listening on Port: " + port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
