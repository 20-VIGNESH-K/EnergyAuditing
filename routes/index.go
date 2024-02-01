package routes

import (
	Controllers "github.com/20-VIGNESH-K/EnergyAuditing/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()


	router.Static("/register", "./frontend/createUser/")

	router.POST("/register", Controllers.CreateUser)
	router.POST("/login", Controllers.Login)

	return router
}
