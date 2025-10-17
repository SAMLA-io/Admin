package vapi

import (
	"errors"
	"net/http"
	"time"

	vapiclient "github.com/VapiAI/server-sdk-go/client"
	"github.com/VapiAI/server-sdk-go/option"
)

func EnsureVapiClient() error {
	if VapiClient == nil {
		return errors.New("VapiClient is not initialized - check VAPI_API_KEY environment variable")
	}
	return nil
}

func CreateClient(apiKey string) *vapiclient.Client {
	return vapiclient.NewClient(
		option.WithToken(apiKey),
		option.WithHTTPClient(
			&http.Client{
				Timeout: 30 * time.Second,
			}),
	)
}
