package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	//connection handler for the DBConnection
	db *sql.DB
)


const (
	sqlStatementInsert string = `
	INSERT INTO "users" (first_name, last_name,birthday, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Data      string
	Email     string
}




const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "theater"
)


func InitConnection()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func CloseDbConnection()  {
	defer db.Close()

}

func CreateTablesIfNotExists()  {
	createUsers :=`CREATE TABLE IF NOT EXISTS "users" (
				  id SERIAL PRIMARY KEY, 
				  first_name TEXT,
				  last_name TEXT,
		          birthday DATE,
		          email TEXT UNIQUE NOT NULL
				  )`

	theater :=`CREATE TABLE IF NOT EXISTS theater (
				  id SERIAL PRIMARY KEY,
				  name TEXT, 
				  rows INT,
				  floor INT,
		          capacity INT
				  )`

	movie :=`CREATE TABLE IF NOT EXISTS movie (
				  id SERIAL PRIMARY KEY, 
				  movie_title TEXT UNIQUE NOT NULL,
				  movie_year INT,
		          movie_time INT
				  )`


	schedule :=`CREATE TABLE IF NOT EXISTS schedule (
				  id SERIAL PRIMARY KEY, 
				  id_theater INT  REFERENCES theater(id),
				  movie_year INT,
		          movie_time INT
				  )`



	booking :=`CREATE TABLE IF NOT EXISTS booking (
				  id SERIAL PRIMARY KEY, 
				  id_person INT  REFERENCES  "user"(id),
				  id_schedule INT REFERENCES schedule(id),
		          place INT
				  )`

	_, err := db.Exec(createUsers)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(movie)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(theater)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(booking)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(schedule)
	if err != nil {
		panic(err)
	}
}

func InsertUser(firstName string, lastName string, date string, email string) (string, error)  {
	InitConnection()
	id := 0
	fmt.Print("Hereee2")
	fmt.Println("parameters are", firstName,lastName,email,date)
	if insert := db.QueryRow(sqlStatementInsert, firstName,  lastName, date, email).Scan(&id) ; insert != nil {
		if pgerr, ok := insert.(*pq.Error); ok {
			if pgerr.Code == "23505" {
				return "", errors.New("user with such an email already exists")
			}
		}
		return "", insert
	}else {

		return fmt.Sprintf("Succesfully inserted %s at index %v", firstName, id), insert
	}
}

func GetUserForId(){
	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 10)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, firstName)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}