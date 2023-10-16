package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/jwt"

	user_domain "server/domain"
	user_services "server/services"
	configs "server/config"
)

type LoginController struct {
	User_DB user_services.UserPort
	Env *configs.Env
}

func (uc *LoginController) Login(c *gin.Context) {

	User := &user_domain.Users{}
	err := c.ShouldBindJSON(User)
	if err != nil {
		// Devuelve una respuesta de error si hay un error al deserializar los datos
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos"})
		return
	}


	if User.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el nombre de usuario en los datos recibidos"})
		return
	}

	if User.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta la contraseña en los datos recibidos"})
		return
	}


	if User.Username == "" && User.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan tanto el nombre de usuario como la contraseña en los datos recibidos"})
		return
	}

	is_correct_info, err := uc.User_DB.Login(User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !is_correct_info{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creando jwt"})
		return
	}
	
	jwtKey, err := jwt.CreateJWT(User, uc.Env)
	// Return an error if the JWT token generation fails.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creando jwt"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token": jwtKey})
}
