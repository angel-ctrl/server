package jwt

import (
	"time"

	configs "server/config"
	user_domain "server/domain"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWT(User *user_domain.Users, config *configs.Env) (string, error) {

	
	// Set the secret key to sign the token
	miClave := []byte(config.SecretKeyJWT)

	// Define the payload with the user data and expiration time
	payload := jwt.MapClaims{
		"name": User.Username,
		"id":   User.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time: 24 hours from now
	}

	// Generate a new token with the payload and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	// Return the generated token or an error if occurred
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
