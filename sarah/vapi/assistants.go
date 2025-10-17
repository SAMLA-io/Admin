package vapi

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VapiAI/server-sdk-go/option"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"

	mongodb "samla-admin/sarah/mongodb"
	mongodbTypes "samla-admin/types/mongodb"

	vapiApi "github.com/VapiAI/server-sdk-go"
	vapiclient "github.com/VapiAI/server-sdk-go/client"
)

var VapiClient *vapiclient.Client

func inti() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[SARAH] WARNING: .env file not found, using system environment variables")
	}

	apiKey := os.Getenv("VAPI_API_KEY")
	if apiKey == "" {
		log.Printf("ERROR: VAPI_API_KEY environment variable is not set")
		return
	}

	VapiClient = CreateClient(apiKey)
	if VapiClient == nil {
		log.Printf("ERROR: Failed to create VapiClient")
	}
}

// ensureVapiClient checks if VapiClient is properly initialized
func ensureVapiClient() error {
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
func CreateAssistant(orgId string, assistantCreateDto *vapiApi.CreateAssistantDto) (*mongo.InsertOneResult, error) {
	if err := ensureVapiClient(); err != nil {
		return nil, err
	}

	assistant, err := VapiClient.Assistants.Create(context.Background(), assistantCreateDto)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := mongodb.CreateAssistant(orgId, mongodbTypes.Assistant{
		Name:            *assistant.Name,
		VapiAssistantId: assistant.Id,
		Type:            "placeholder type",
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
