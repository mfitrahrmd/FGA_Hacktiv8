package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 1")

		next.ServeHTTP(w, r)
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 2")

		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", middleware1(middleware2(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})))

	log.Fatalln(http.ListenAndServe(":80", nil))
}
