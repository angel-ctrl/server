package routers

import (
	"server/controllers"

	"github.com/gin-gonic/gin"

	"server/middlewares"
)

func (d *HttpSetup) NewLoginRouter(group *gin.RouterGroup) {

	uc := controllers.LoginController{
		User_DB:  d.UserDB,
		Env: d.Env,
	}

	group.POST("/login/v1", middlewares.AuthMiddleware(d.UserDB), uc.Login)
}
