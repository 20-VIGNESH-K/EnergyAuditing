package routes

import (
	Controllers "github.com/20-VIGNESH-K/EnergyAuditing/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Serve static files for the register and login pages
	router.Static("/register", "./frontend/createUser/")
	router.Static("/profile", "./frontend/profile")
	router.Static("/login", "./frontend/login/")
	router.Static("/logout", "./frontend/logout/")

	router.Static("/home", "./frontend/home/")
	router.Static("/home-industry", "./frontend/industry")
	router.Static("/listindustry", "./frontend/dropdown")
	
	router.Static("/weaving", "./frontend/industry/weaving")
	router.Static("/textile", "./frontend/industry/textile")
	router.Static("/it", "./frontend/industry/it")

	router.Static("/result", "./frontend/result/")
	

	// Handle POST requests for user registration and login
	router.POST("/register", Controllers.CreateUser)
	router.POST("/login", Controllers.Login)
	router.POST("/weaving", Controllers.Weaving)
	router.POST("/textile", Controllers.Textile)
	router.POST("/it", Controllers.IT)
	router.POST("/getuser",Controllers.GetUser)
	return router
}
