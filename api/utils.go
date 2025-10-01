package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/organization"
)

func VerifyMethod(r *http.Request, allowedMethods []string) bool {
	for _, method := range allowedMethods {
		if r.Method == strings.ToUpper(method) {
			return true
		}
	}
	return false
}

func ExtractOrganizationCreateRequest(r *http.Request) (organization.CreateParams, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return organization.CreateParams{}, err
	}

	var organization organization.CreateParams
	if err := json.Unmarshal(body, &organization); err != nil {
		return organization, err
	}

	return organization, nil
}

func ExtractOrganizationUpdateRequest(r *http.Request) (organization.UpdateParams, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return organization.UpdateParams{}, err
	}

	var updateRequest organization.UpdateParams
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		return updateRequest, err
	}

	return updateRequest, nil
}

func ExtractOrganizationId(r *http.Request) (string, error) {
	organizationId := r.URL.Query().Get("organization_id")
	if organizationId == "" {
		return "", errors.New("organization_id is required")
	}
	return strings.TrimSpace(organizationId), nil
}
