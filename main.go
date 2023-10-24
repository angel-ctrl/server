package main

import (
	configs "server/config"
	database "server/database"
	user_online "server/online"
	"server/routers"
	user_services "server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	Env := configs.CreateConfigs()

	DB := database.SetupDatabaseConnection(Env)

	user_services_var := user_services.CreateUsersServicesImplementation(DB)

	Hub := user_online.NewHub(Env, user_services_var)
	go Hub.UsersManager()

	router := gin.Default()

	cors_server := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	})

	router.Use(cors_server)

	httpService := routers.NewHttpSetup(router, Env, user_services_var, Hub)
	httpService.Setup()

	router.Run(Env.Host + ":" + Env.Port)

}
