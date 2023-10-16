package services

import (
	"server/utils"
	user_domain "server/domain"

	"github.com/google/uuid"
)

func (serviceResources *Users_services_implmentation) CreateUser(UserData *user_domain.Users) (*user_domain.Users, error) {
	id := uuid.New()

	// Encrypt the user password
	pass, err := utils.EncriptPass(UserData.Password)
	if err != nil {
		return nil, err
	}

	// Execute the insert query with the user data
	_, err = serviceResources.DB.Exec(`
		INSERT INTO users (
			id,
			username,
			password,
			user_state
		)
		VALUES ($1, $2, $3, $4)
	`, id, UserData.Username, pass, UserData.UserState)
	if err != nil {
		return nil, err
	}

	// Clear the password field for security reasons
	UserData.Password = ""

	UserData.ID = id.String()

	// Return the created user
	return UserData, nil
}
