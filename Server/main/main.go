package main

import (
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Server/database"
)


func main() {



	database.InitConnection()
	database.CreateTablesIfNotExists()
	//
	//gin.SetMode(gin.ReleaseMode)
	router := Middleware.SetupRouter()
	router.Run(":8080")


	//database.CloseDbConnection()

}




