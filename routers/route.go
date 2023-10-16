package routers

import (
	"server/config"

	user_services "server/services"
	user_online "server/online"

	"github.com/gin-gonic/gin"
)

type HttpSetup struct {
	Gin    *gin.Engine
	Env    *config.Env
	UserDB user_services.UserPort
}

func NewHttpSetup(gin *gin.Engine, env *config.Env, user_DB user_services.UserPort) *HttpSetup {
	return &HttpSetup{
		Gin:    gin,
		Env:    env,
		UserDB: user_DB,
	}
}

func (d *HttpSetup) Setup(hub *user_online.Hub) {

	publicRouter := d.Gin.Group("API")
	// All Public APIs
	d.NewUserCreationRouter(publicRouter)
	//d.NewGetUserRouter(publicRouter)
	d.NewLoginRouter(publicRouter)
	d.NewUpdatePlanRouter(publicRouter)
	d.NewBanUserRouter(publicRouter)
	publicRouter.GET("/online/v1", hub.HandlerWebSocket)

	/*
		protectedRouter := gin.Group("")
		// Middleware to verify AccessToken
		protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
		// All Private APIs
		NewProfileRouter(env, timeout, db, protectedRouter)
		NewTaskRouter(env, timeout, db, protectedRouter)*/
}
