package controllers

import (
	"net/http"

	user_domain "server/domain"

	user_services "server/services"

	user_online "server/online"

	"github.com/gin-gonic/gin"
)

type BanUserController struct {
	User_DB user_services.UserPort
	Hub *user_online.Hub
}

func (uc *BanUserController) BanUser(c *gin.Context) {
	// Obtén los parámetros de la solicitud
	User := &user_domain.Users{}
	err := c.ShouldBindJSON(User)
	if err != nil {
		// Devuelve una respuesta de error si hay un error al deserializar los datos
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos"})
		return
	}

	err = User.Validate_without_pass()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	if User.UserState != "ban" && User.UserState != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	// Realiza la actualización del plan
	err = uc.User_DB.BanUser(User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el plan"})
		return
	}

	uc.Hub.DisconnectUser(User.Username)

	c.JSON(http.StatusOK, gin.H{"message": "user actualizado exitosamente"})
}
