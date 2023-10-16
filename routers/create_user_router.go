package routers

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func (d *HttpSetup) NewUserCreationRouter(group *gin.RouterGroup) {

	uc := controllers.UserCreationController{
		User_DB: d.UserDB,
		Env:     d.Env,
	}

	group.POST("/create_user/v1", middlewares.SuperPassValidateMiddleware(d.Env), uc.UserCreation)
}
