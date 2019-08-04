package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/instructions", app.CreateUser)
		v1.POST("/create_user", app.CreateUser)
		v1.PUT("/instructions/:id", app.UpdateInstruction)
		v1.DELETE("/instructions/:id", app.DeleteInstruction)
	}

	return router
}
