package controllers

import (
	"net/http"
	"strconv"

	user_domain "server/domain"

	user_services "server/services"

	"github.com/gin-gonic/gin"
)

type UpdatePlanController struct {
	User_DB user_services.UserPort
}

func (uc *UpdatePlanController) UpdatePlan(c *gin.Context) {

    User := &user_domain.Users{}
	err := c.ShouldBindJSON(User)
	if err != nil {
		// Devuelve una respuesta de error si hay un error al deserializar los datos
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos"})
		return
	}

	err = User.Validate_without_pass_userstate()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

    daysToAddStr := c.Query("daysToAdd")

    // Verifica si se proporcionaron los parámetros necesarios
    if daysToAddStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan parámetros requeridos"})
        return
    }

    // Convierte los días a agregar a un entero
    daysToAdd, err := strconv.Atoi(daysToAddStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Los días a agregar no son válidos"})
        return
    }

    // Obten el usuario de la base de datos
    user, err := uc.User_DB.GetUser(User)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario"})
        return
    }

    // Realiza la actualización del plan
    err = uc.User_DB.UpdatePlan(user, daysToAdd)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el plan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Plan actualizado exitosamente"})
}