package routers

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func (d *HttpSetup) NewUpdatePlanRouter(group *gin.RouterGroup) {

	uc := controllers.UpdatePlanController{
		User_DB: d.UserDB,
	}

	group.PATCH("/update_plan/v1", middlewares.SuperPassValidateMiddleware(d.Env), uc.UpdatePlan)
}
