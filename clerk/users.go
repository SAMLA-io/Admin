package clerk

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

func GetAllUsers() (*clerk.UserList, error) {
	users, err := user.List(context.Background(), &user.ListParams{})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(userId string) (*clerk.User, error) {
	user, err := user.Get(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(createParams *user.CreateParams) (*clerk.User, error) {
	user, err := user.Create(context.Background(), createParams)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(userId string, updateParams *user.UpdateParams) (*clerk.User, error) {
	user, err := user.Update(context.Background(), userId, updateParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userId string) (*clerk.DeletedResource, error) {
	resource, err := user.Delete(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
