package clerk

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/organization"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
}

func GetAllOrganizations() (*clerk.OrganizationList, error) {
	organizations, err := organization.List(context.Background(), &organization.ListParams{})
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func GetOrganizationUsers(organizationId string) (*clerk.UserList, error) {
	users, err := user.List(context.Background(), &user.ListParams{
		OrganizationIDs: []string{organizationId},
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserOrganizations(userId string) (*clerk.OrganizationMembershipList, error) {
	orgMemberships, err := user.ListOrganizationMemberships(context.Background(), userId, &user.ListOrganizationMembershipsParams{})

	if err != nil {
		log.Printf("Error getting organization memberships: %v", err)
	}

	return orgMemberships, nil
}

func GetUserOrganizationId(userId string) (string, error) {
	orgMemberships, err := GetUserOrganizations(userId)
	if err != nil {
		return "", err
	}
	if len(orgMemberships.OrganizationMemberships) == 0 {
		return "", errors.New("no organization memberships found")
	}
	return orgMemberships.OrganizationMemberships[0].Organization.ID, nil
}

func GetOrganizationPublicMetadata(organizationId string) (map[string]interface{}, error) {
	organization, err := organization.Get(context.Background(), organizationId)
	if err != nil {
		return nil, err
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal(organization.PublicMetadata, &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

func UpdateOrganizationPublicMetadata(organizationId string, metadata map[string]interface{}) error {
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	rawMessage := json.RawMessage(jsonData)
	_, err = organization.Update(context.Background(), organizationId, &organization.UpdateParams{
		PublicMetadata: &rawMessage,
	})

	return err
}

func UpdateOrganization(organizationId string, updateRequest organization.UpdateParams) (clerk.Organization, error) {
	organization, err := organization.Update(context.Background(), organizationId, &updateRequest)

	if err != nil {
		return *organization, err
	}

	return *organization, nil
}

func CreateOrganization(newOrganization organization.CreateParams) (clerk.Organization, error) {
	organization, err := organization.Create(context.Background(), &newOrganization)
	if err != nil {
		return clerk.Organization{}, err
	}

	return *organization, nil
}

func DeleteOrganization(organizationId string) error {
	_, err := organization.Delete(context.Background(), organizationId)
	if err != nil {
		return err
	}

	return nil
}
