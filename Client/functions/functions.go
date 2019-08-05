package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olekukonko/tablewriter"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/app"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

const (
	Host     = "localhost"
	Port     =  8080
	Group    = "/api/v1"
)


type Movie struct {
	ID           int
	Movie_title  string
	Movie_year   int
	Pg_type      string
	Runtime      int
}

var Router *gin.Engine

func CreateUser(first_name, last_name, birthday, email string) error {

	m := app.User{FirstName:first_name,  LastName:last_name, Birthday:"10-16-2020", Email: email}
	b, err := json.Marshal(m)


	req, err := http.NewRequest("POST", "/api/v1/create_user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Post hearteat failed with error %d.", err)
		return err
	}

	resp := httptest.NewRecorder()
	Router.ServeHTTP(resp, req)

	if resp.Code != 200 {
		fmt.Println("/api/v1/instructions failed with error code %d and response", resp.Code, resp.Body)
	}else
	{
		fmt.Println(resp.Body)
	}
	return nil
}

func AddNewTheater(name, rows, floor string) error {

	m := app.Theater{Name:name,  Rows:rows, Floor: floor}
	b, err := json.Marshal(m)

	resp, err := http.Post("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/add_theater", "application/json", bytes.NewBuffer(b))

	if err != nil {
		fmt.Printf("Post request failed for creating new theater with error %d.", err)
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("/api/v1/add_movie failed with error code %s and response %s", resp, resp.Body)
	}else{
		fmt.Println("Success!")
	}
	return nil

}

func GetAvailableMovies() error {
	req, err := http.NewRequest("GET", "/api/v1/movies",nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Post hearteat failed with error %d.", err)
		return err
	}

	resp := httptest.NewRecorder()
	Router.ServeHTTP(resp, req)
	if resp.Code != 200 {
		fmt.Println("/api/v1/movies failed with error code %d and response", resp.Code, resp.Body)
	}else
	{
		movies := []Movie{}
		data := [][]string{}
		err := json.Unmarshal(resp.Body.Bytes(),&movies)
		if err != nil{
			fmt.Printf(err.Error())
		}
		for _, movie := range movies {
			data = append(data, []string{ strconv.Itoa(movie.ID),movie.Movie_title, strconv.Itoa(movie.Movie_year), movie.Pg_type, strconv.Itoa(movie.Runtime)})
		}


		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID","Movie Title", "Movie Year", "Pg_type", "Runtime"})
		table.SetFooter([]string{"Total Movies: " + strconv.Itoa(len(movies)), "", "", "",""}) // Add Footer
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetBorder(false)                                // Set Border to false
		table.AppendBulk(data)                                // Add Bulk Data
		table.Render()


	}
	return nil
}

func InitRouter (){
	Router = Middleware.SetupRouter()
}
