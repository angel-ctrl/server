package routers

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func (d *HttpSetup) NewGetUserRouter(group *gin.RouterGroup) {

	uc := controllers.UserGetController{
		User_DB:  d.UserDB,
	}

	group.GET("/get_user/v1", uc.GetUser)
}
