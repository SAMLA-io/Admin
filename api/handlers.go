package api

import (
	"encoding/json"
	"net/http"
	"samla-admin/clerk"
)

func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"GET"}) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	organizations, err := clerk.GetAllOrganizations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(organizations)
}
