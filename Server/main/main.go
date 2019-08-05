package main

import (
	"fmt"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Server/database"
)


func main() {

	database.InitConnection()
	database.CreateTablesIfNotExists()
	database.RunProvisioning(false)

	_, err :=database.ListMovies()
	if err != nil{
		fmt.Println("failed")
		fmt.Println(err)
	}else{
		fmt.Println("success")
	}
	//
	//gin.SetMode(gin.ReleaseMode)
	router := Middleware.SetupRouter()
	router.Run("localhost:8080")



}




