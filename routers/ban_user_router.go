package routers

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func (d *HttpSetup) NewBanUserRouter(group *gin.RouterGroup) {

	uc := controllers.BanUserController{
		User_DB: d.UserDB,
		Hub:     d.Hub,
	}

	group.PATCH("/ban_user/v1", middlewares.SuperPassValidateMiddleware(d.Env), uc.BanUser)
}
