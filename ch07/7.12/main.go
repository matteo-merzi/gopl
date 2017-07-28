package main

import (
	"net/http"

	"github.com/matteo-merzi/gopl/ch07/7.12/database"
)

func main() {
	m := &database.MutexDB{}
	http.HandleFunc("/list", m.List)
	http.HandleFunc("/price", m.Price)
	http.HandleFunc("/create", m.Create)
	http.HandleFunc("/read", m.Read)
	http.HandleFunc("/update", m.Update)
	http.HandleFunc("/delete", m.Delete)
	http.ListenAndServe("localhost:8080", nil)
}
