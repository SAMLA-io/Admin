package vapi

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"

	mongodb "samla-admin/sarah/mongodb"
	mongodbTypes "samla-admin/types/mongodb"

	vapiApi "github.com/VapiAI/server-sdk-go"
	vapiclient "github.com/VapiAI/server-sdk-go/client"
)

var VapiClient *vapiclient.Client

func init() {
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

func CreateAssistant(orgId string, assistantCreateDto *vapiApi.CreateAssistantDto) (*mongo.InsertOneResult, error) {
	if err := EnsureVapiClient(); err != nil {
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

func UpdateAssistant(assistantId string, assistantUpdateDto vapiApi.UpdateAssistantDto) (*vapiApi.Assistant, error) {
	if err := EnsureVapiClient(); err != nil {
		return nil, err
	}

	result, err := VapiClient.Assistants.Update(context.Background(), assistantId, &assistantUpdateDto)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func DeleteAssistant(orgId string, assistantId string) (*mongo.DeleteResult, error) {
	if err := EnsureVapiClient(); err != nil {
		return nil, err
	}

	_, err := VapiClient.Assistants.Delete(context.Background(), assistantId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := mongodb.DeleteAssistant(orgId, assistantId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func ExistsAssistant(assistantId string) bool {
	if err := EnsureVapiClient(); err != nil {
		log.Printf("ExistsAssistant: VapiClient not initialized: %v", err)
		return false
	}

	log.Printf("ExistsAssistant: Checking existence of assistant with ID: %s", assistantId)
	assistant, err := VapiClient.Assistants.Get(context.Background(), assistantId)
	if err != nil {
		log.Printf("ExistsAssistant: Error fetching assistant with ID %s: %v", assistantId, err)
		return false
	}

	exists := assistant != nil && assistant.Id != ""
	log.Printf("ExistsAssistant: Assistant with ID %s exists: %v", assistantId, exists)
	return exists
}
