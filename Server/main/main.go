package main

import (
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Server/database"
)


func main() {

	database.InitConnection()
	database.CreateTablesIfNotExists()
	database.RunProvisioning(false)
	//
	//gin.SetMode(gin.ReleaseMode)
	router := Middleware.SetupRouter()
	router.Run("localhost:8080")



}




