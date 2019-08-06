package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Server/database"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/structs"
	"net/http"
)


func CreateUser(c *gin.Context) {
	dataRequest, _ := c.GetRawData()

	var data *structs.User
	err := json.Unmarshal(dataRequest,&data)
	if err := json.Unmarshal(dataRequest,data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := database.InsertUser(data.FirstName,data.LastName, data.Birthday, data.Email)
	if err == nil {
		c.JSON(201, content)
	} else {
		c.JSON(500, err.Error())
	}


}

func AddMovie(c *gin.Context)  {
	request,_ := c.GetRawData()

	var data *structs.NewMovie

	err := json.Unmarshal(request,&data)
	if err := json.Unmarshal(request,data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := database.AddMovie(data)
	fmt.Println("Error",response, err)
	if err == nil {

		c.JSON(201, response)
	} else {
		c.JSON(500, err.Error())
	}

}

func AddTheater(c *gin.Context) {
	dataRequest, _ := c.GetRawData()

	var data *structs.Theater
	err := json.Unmarshal(dataRequest,&data)
	if err := json.Unmarshal(dataRequest,data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    //refactor to pass the whole structure
	response, err := database.AddTheater(data.Name,data.Rows, data.Floor)
	if err == nil {

		c.JSON(201, response)
	} else {
		c.JSON(500, err.Error())
	}
}

func ListAllMovies(c *gin.Context){
	content, err := database.ListMovies()
	if err == nil {
		c.JSON(200, content)
	} else {
		c.JSON(500, err.Error())
	}
}

func ListAllTheaters(c *gin.Context){
	content, err := database.ListTheaters()
	if err == nil {
		c.JSON(200, content)
	} else {
		c.JSON(500, err.Error())
	}
}

func PostInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
}

func UpdateInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
}

func DeleteInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}