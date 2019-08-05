package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Server/database"
	"net/http"
)

// Bindings from and to JSON
type User struct {
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
}

// Bindings from and to JSON
type Theater struct {
	Name    string
	Rows     string
	Floor     string
}



func GetUserForId(c *gin.Context) {

	//err, content := database.GetUserForId()
	//if err == nil {
	//
	//	c.JSON(200, content)
	//} else {
	//	c.JSON(404, gin.H{"error": "instruction not found"})
	//}
	c.JSON(404, gin.H{"error": "instruction not found"})


}

func CreateUser(c *gin.Context) {
	dataRequest, _ := c.GetRawData()

	var data *User
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

func AddTheater(c *gin.Context) {
	dataRequest, _ := c.GetRawData()

	var data *Theater
	err := json.Unmarshal(dataRequest,&data)
	if err := json.Unmarshal(dataRequest,data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

func PostInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})

}

func UpdateInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})

}

func DeleteInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}