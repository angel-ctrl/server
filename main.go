package main

import (
	configs "server/config"
	database "server/database"
	user_online "server/online"
	"server/routers"
	user_services "server/services"

	"github.com/gin-gonic/gin"
)

func main(){

	Env := configs.CreateConfigs()

	DB := database.SetupDatabaseConnection(Env)

	user_services_var := user_services.CreateUsersServicesImplementation(DB)

	Hub := user_online.NewHub(Env, user_services_var)
	go Hub.UsersManager()

	gin := gin.Default()
	httpService := routers.NewHttpSetup(gin, Env, user_services_var)
	httpService.Setup(Hub)

	gin.Run(Env.Host + ":" + Env.Port)

}