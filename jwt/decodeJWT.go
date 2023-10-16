package jwt

import (
    "errors"
    "server/config"
    "strings"
    "fmt"

    jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
    Username string `json:"name"`
    Id       string `json:"id"`
    jwt.StandardClaims
}

func DecodeToken(tk string, Env *config.Env) (*Claim, bool, error) {
    miClave := []byte(Env.SecretKeyJWT)

    claims := &Claim{}

    splitToken := strings.Split(tk, "Bearer")

    if len(splitToken) != 2 {
        fmt.Println("Formato de token inválido")
        return claims, false, errors.New("formato de token inválido")
    }

    tk = strings.TrimSpace(splitToken[1])

    tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
        return miClave, nil
    })

    if err != nil {
        fmt.Println("Error al decodificar token:", err)
        return claims, false, err
    }

    if !tkn.Valid {
        fmt.Println("Token inválido")
        return claims, false, errors.New("token inválido")
    }

    return claims, true, nil
}


func ExtractClaims(tokenStr string, Env *config.Env) (jwt.MapClaims, bool) {
	hmacSecret := []byte(Env.SecretKeyJWT)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		 // check token signing method etc
		 return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}