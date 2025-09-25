package main

import (
	"fmt"
	"log"
	"net/http"
	"samla-admin/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Samla Admin!")
	})

	http.HandleFunc("/organizations", api.GetAllOrganizations) // GET: All organizations

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
