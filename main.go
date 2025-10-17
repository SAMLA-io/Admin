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

	/* General ADMIN Endpoints */

	http.Handle("/organizations/all", auth.VerifyingMiddleware(http.HandlerFunc(api.GetAllOrganizations)))    // GET: All organizations
	http.Handle("/organizations/create", auth.VerifyingMiddleware(http.HandlerFunc(api.CreateOrganization)))  // POST: Create organization
	http.Handle("/organizations/update", auth.VerifyingMiddleware(http.HandlerFunc(api.UpdateOrganization)))  // PATCH: Update organization
	http.Handle("/organizations/delete", auth.VerifyingMiddleware(http.HandlerFunc(api.DeleteOrganization)))  // DELETE: Delete rganization
	http.Handle("/organizations/users", auth.VerifyingMiddleware(http.HandlerFunc(api.GetOrganizationUsers))) // GET: Get organization users

	http.Handle("/users/all", auth.VerifyingMiddleware(http.HandlerFunc(api.GetAllUsers)))   // GET: All users
	http.Handle("/users/get", auth.VerifyingMiddleware(http.HandlerFunc(api.GetUser)))       // GET: Get user
	http.Handle("/users/create", auth.VerifyingMiddleware(http.HandlerFunc(api.CreateUser))) // POST: Create user
	http.Handle("/users/update", auth.VerifyingMiddleware(http.HandlerFunc(api.UpdateUser))) // PATCH: Update user
	http.Handle("/users/delete", auth.VerifyingMiddleware(http.HandlerFunc(api.DeleteUser))) // DELETE: Delete user

	http.Handle("/invitations/all", auth.VerifyingMiddleware(http.HandlerFunc(api.GetAllInvitations)))   // GET: All invitations
	http.Handle("/invitations/create", auth.VerifyingMiddleware(http.HandlerFunc(api.CreateInvitation))) // POST: Create invitation

	/* Sarah Admin Endpoints */
	http.Handle("/sarah/assistants/organization", auth.VerifyingMiddleware(http.HandlerFunc(api.GetOrganizationAssistants))) // GET: Get organization assistants
	http.Handle("/sarah/assistants/create", auth.VerifyingMiddleware(http.HandlerFunc(api.CreateAssistant)))                 // POST: Create assistant

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
