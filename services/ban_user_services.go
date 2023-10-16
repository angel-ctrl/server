package services

import (
	user_domain "server/domain"
)

func (serviceResources *Users_services_implmentation) BanUser(User *user_domain.Users) error {

	// Realiza la actualización del estado del usuario en la base de datos
	_, err := serviceResources.DB.Exec(`
        UPDATE users
        SET user_state = $1
        WHERE username = $2
    `, User.UserState, User.Username)
	if err != nil {
		return err
	}

	return nil
}
