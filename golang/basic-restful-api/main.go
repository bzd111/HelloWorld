package main

import (
	"basic-restful-api/handlers"
	"basic-restful-api/storage"
	"log"
	"net/http"
)

func main() {
	db := storage.NewInMemoryDB()
	mux := http.NewServeMux()
	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))
	log.Printf("serving on port 8080")

	// err := http.List
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
