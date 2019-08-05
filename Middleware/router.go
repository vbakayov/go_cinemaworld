package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/movies", app.ListAllMovies)
		v1.GET("/theaters", app.ListAllTheaters)
		v1.POST("/create_user", app.CreateUser)
		v1.POST("/add_theater", app.AddTheater)
		v1.DELETE("/instructions/:id", app.DeleteInstruction)
	}

	return router
}
