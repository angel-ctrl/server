package services

import (
	"errors"
	user_domain "server/domain"

	"golang.org/x/crypto/bcrypt"
)

func (serviceResources *Users_services_implmentation) Login(UserData *user_domain.Users) (bool, error) {

	UserLooked, err := serviceResources.GetUser(&user_domain.Users{Username: UserData.Username})

	if err != nil {
		return false, errors.New("usuario no encontrado")
	}

	passwordBytes := []byte(UserData.Password)
	passwordBD := []byte(UserLooked.Password)
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return false, errors.New("contraseña incorrecta")
	}

	return true, nil
}
