package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/structs"
	"io/ioutil"
	"net/http"
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


func CreateUser(first_name, last_name, birthday, email string) error {

	m := structs.User{FirstName:first_name,  LastName:last_name, Birthday:birthday, Email: email}
	b, err := json.Marshal(m)

	resp, err := http.Post("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/create_user", "application/json", bytes.NewBuffer(b))


	if err != nil {
		fmt.Printf("Post request failed for creating new user with error %d.", err)
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if  resp.StatusCode != http.StatusCreated {
		fmt.Printf("/api/v1/instructions failed with error code %d and response  %s", resp.StatusCode, body)
	}else
	{
		fmt.Println("Success!")
	}
	return nil
}

func AddNewTheater(name, rows, floor string) error {

	m := structs.Theater{Name:name,  Rows:rows, Floor: floor}
	b, err := json.Marshal(m)

	resp, err := http.Post("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/add_theater", "application/json", bytes.NewBuffer(b))

	if err != nil {
		fmt.Printf("Post request failed for creating new theater with error %d.", err)
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("/api/v1/add_movie failed with error code %d and response %s", resp.StatusCode, body)
	}else{
		fmt.Println("Success!")
	}
	return nil

}

func AddMovie(name string, movieYear string, pgType string, runtime string, theater string, schedule map[string][]string) error  {
	m := structs.NewMovie{Name:name, MovieYear: movieYear, PgType: pgType,Runtime:runtime, Theater:theater, Schedule:schedule}
	b, _ := json.Marshal(m)

	resp, err := http.Post("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/add_movie", "application/json", bytes.NewBuffer(b))

	if err != nil {
		fmt.Printf("Post request failed for adding a new movie with error %d.", err)
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("/api/v1/add_movie failed with error code %d and response %s", resp.StatusCode, body)
	}else{
		fmt.Printf("Success!  code %d and response %s", resp.StatusCode,body)
	}
	return nil

}

func GetAvailableTheaters() ([]string,error) {
	data := []string{}
	resp, err := http.Get("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/theaters")

	if err != nil {
		fmt.Printf("Get request failed for listing the available theaters with error %d.", err)
		return nil,err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("/api/v1/theaters failed with error code %d and response %s", resp.StatusCode, body)
		return nil, errors.New("response code not as expected 201")

	}else
	{
		theaters := []structs.Theater{}
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body,&theaters)
		if err != nil{
			fmt.Printf(err.Error())
		}
		for _, theater := range theaters {
			data = append(data, theater.Name +" PgType: " + theater.Floor +" MovieYear: " + theater.Rows)
		}


	}
	return data,nil



}

func GetAvailableMovies() error {

	resp, err := http.Get("http://"+ Host +":"+ strconv.Itoa(Port) + Group + "/movies")

	if err != nil {
		fmt.Printf("Get request failed for listing the available movies with error %d.", err)
		return err
	}


	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("/api/v1/movies failed with error code %d and response %s", resp.StatusCode, body)
	}else
	{
		movies := []Movie{}
		data := [][]string{}
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body,&movies)
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
