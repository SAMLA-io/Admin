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

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"POST"}) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	organization, err := ExtractOrganizationCreateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrganization, err := clerk.CreateOrganization(organization)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdOrganization)

}

func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"PATCH"}) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	organizationId, err := ExtractOrganizationId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateRequest, err := ExtractOrganizationUpdateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedOrganization, err := clerk.UpdateOrganization(organizationId, updateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedOrganization)
}
