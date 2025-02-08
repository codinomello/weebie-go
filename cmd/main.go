package main

import (
	"net/http"

	"github.com/codinomello/webjetos-go/db"
)

func main() {
	db.DBConnection()
	server := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/hello", server)
	http.ListenAndServe(":8080", server)
}
