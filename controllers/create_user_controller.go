package controllers

import (
	"net/http"

	configs "server/config"
	user_domain "server/domain"
	user_services "server/services"

	"github.com/gin-gonic/gin"
)

type UserCreationController struct {
	User_DB user_services.UserPort
	Env     *configs.Env
}

func (uc *UserCreationController) UserCreation(c *gin.Context) {

	User := &user_domain.Users{}
	err := c.ShouldBindJSON(User)
	if err != nil {
		// Devuelve una respuesta de error si hay un error al deserializar los datos
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos"})
		return
	}

	err = User.Validate()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	user, err := uc.User_DB.CreateUser(User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Clear the password field for security reasons
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
