package main

import (
	"fmt"
	"log"
	"net/http"
	"samla-admin/api"
	"samla-admin/auth"
	"time"
)

func main() {
	http.HandleFunc("/", welcome)

	http.Handle("/organizations/all", auth.TestingMiddleware(http.HandlerFunc(api.GetAllOrganizations)))   // GET: All organizations
	http.Handle("/organizations/create", auth.TestingMiddleware(http.HandlerFunc(api.CreateOrganization))) // POST: Create organization
	http.Handle("/organizations/update", auth.TestingMiddleware(http.HandlerFunc(api.UpdateOrganization))) // PATCH: Update organization
	http.Handle("/organizations/delete", auth.TestingMiddleware(http.HandlerFunc(api.DeleteOrganization))) // DELETE: Delete organization

	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Starting Samla Admin API on port 8080...\n")
	log.Fatal(server.ListenAndServe())
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Sarah AI Call Assistant!"))
}
