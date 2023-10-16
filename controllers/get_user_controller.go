package controllers

import (
	"fmt"
	"net/http"

	user_domain "server/domain"

	user_services "server/services"

	"github.com/gin-gonic/gin"
)

type UserGetController struct {
	User_DB user_services.UserPort
}

func (uc *UserGetController) GetUser(c *gin.Context) {

	User := &user_domain.Users{}

	User.Username = c.Query("username")
	
	user, err := uc.User_DB.GetUser(User)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Clear the password field for security reasons
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
