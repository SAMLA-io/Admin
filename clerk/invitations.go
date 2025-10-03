package clerk

import (
	"context"
	"log"

	"github.com/clerk/clerk-sdk-go/v2"
	organizationInvitation "github.com/clerk/clerk-sdk-go/v2/invitation"
)

func GetAllInvitations() (*clerk.InvitationList, error) {
	invitations, err := organizationInvitation.List(context.Background(), &organizationInvitation.ListParams{})
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func CreateInvitation(userId string) (*clerk.Invitation, error) {
	user, err := GetUser(userId)
	if err != nil {
		return nil, err
	}

	invitationParams := organizationInvitation.CreateParams{
		EmailAddress: user.EmailAddresses[0].EmailAddress,
	}

	invitation, err := organizationInvitation.Create(context.Background(), &invitationParams)
	if err != nil {
		log.Printf("Error creating invitation: %v", err)
		return nil, err
	}
	return invitation, nil
}
