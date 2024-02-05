package routes

import (
	Controllers "github.com/20-VIGNESH-K/EnergyAuditing/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Serve static files for the register and login pages
	router.Static("/register", "./frontend/createUser/")
	router.Static("/login", "./frontend/login/")
	router.Static("/home", "./frontend/home/")
	router.Static("/logout", "./frontend/logout/")

	// Handle POST requests for user registration and login
	router.POST("/register", Controllers.CreateUser)
	router.POST("/login", Controllers.Login)

	return router
}
