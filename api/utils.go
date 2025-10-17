package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	vapiApi "github.com/VapiAI/server-sdk-go"
	"github.com/clerk/clerk-sdk-go/v2/organization"
	"github.com/clerk/clerk-sdk-go/v2/user"
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

func ExtractUserId(r *http.Request) (string, error) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		return "", errors.New("user_id is required")
	}
	return strings.TrimSpace(userId), nil
}

func ExtractUserCreateRequest(r *http.Request) (user.CreateParams, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return user.CreateParams{}, err
	}

	var createRequest user.CreateParams
	if err := json.Unmarshal(body, &createRequest); err != nil {
		return createRequest, err
	}

	return createRequest, nil
}

func ExtractUserUpdateRequest(r *http.Request) (user.UpdateParams, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return user.UpdateParams{}, err
	}

	var updateRequest user.UpdateParams
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		return updateRequest, err
	}
	return updateRequest, nil
}

func ExtractAssistantCreateDto(r *http.Request) *vapiApi.CreateAssistantDto {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	if len(body) == 0 {
		return nil
	}

	var requestBody struct {
		AssistantCreateRequest vapiApi.CreateAssistantDto `json:"assistantCreateRequest"`
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return nil
	}

	return &requestBody.AssistantCreateRequest
}
