package router

import (
	"log"
	"microservice/configuration"
	"microservice/controllers"
	"microservice/executor"
	"microservice/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	City  string `json:"city"`
}

type Users []User

func SetRoutes() *gin.Engine {
	routes := gin.Default()
	_, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalln("unable to load config")
	}

	testExecutor := executor.NewTestExecutor()
	testService := service.NewTestService(testExecutor)
	lController := controllers.NewLoginAuthController(testService)

	routes.GET("service/price", lController.FetchDataController)

	return routes
}
