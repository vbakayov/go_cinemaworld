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
	var query [9]*string

	createMovieType :=	`do 
                         $$
                         begin         
						   IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'screening_type') THEN
					         CREATE TYPE screening_type AS ENUM ('2D', '3D', '4D');
                           END IF;
                         end
						 $$`
	query[0] = &createMovieType

	createPGType :=	`do 
                     $$
                     begin         
                       IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'pg_type_movie') THEN
                         CREATE TYPE pg_type_movie AS ENUM ('G', '12', '15','18');
                       END IF;
                     end
					 $$`
	query[1] = &createPGType

	createUsers :=`CREATE TABLE IF NOT EXISTS "users" (
				  id SERIAL PRIMARY KEY, 
				  first_name TEXT,
				  last_name TEXT,
		          birthday DATE,
		          email TEXT UNIQUE NOT NULL
				  )`
	query[2] = &createUsers

	theater :=`CREATE TABLE IF NOT EXISTS theater (
				  id SERIAL PRIMARY KEY,
				  name TEXT, 
				  rows INT,
				  floor INT,
		          capacity INT
				  )`
	query[3] = &theater

	movie :=`CREATE TABLE IF NOT EXISTS movie (
				  id SERIAL PRIMARY KEY, 
				  movie_title TEXT UNIQUE NOT NULL,
				  movie_year INT,
				  pg_type pg_type_movie,
		          runtime INT
				  )`
	query[4] = &movie

	createSeats :=`CREATE TABLE IF NOT EXISTS seat (
				  id         SERIAL PRIMARY KEY, 
				  theater_id INT  REFERENCES theater(id),
				  row        INT,
				  number     INT
				  )`
	query[5] = &createSeats

	createReservation :=`CREATE TABLE IF NOT EXISTS reservation (
				  id             SERIAL PRIMARY KEY, 
				  screeing_id    INT REFERENCES screening(id),
				  user_id        INT REFERENCES users(id),
				  reserved       BOOLEAN,
				  paid           BOOLEAN,
				  active         BOOLEAN
				  )`
	query[6] = &createReservation

	createReservedSeats :=`CREATE TABLE IF NOT EXISTS reserve_seats (
				  id             SERIAL PRIMARY KEY, 
				  seat_id        INT REFERENCES seat(id),
				  reservation_id INT REFERENCES reservation(id),
				  screeing_id    INT REFERENCES screening(id)
				  )`
	query[7] = &createReservedSeats


	screening :=`CREATE TABLE IF NOT EXISTS screening (
				  id              SERIAL PRIMARY KEY, 
				  id_movie        INT  REFERENCES movie(id),
				  id_theater      INT  REFERENCES theater(id),
				  date            date,    --could be timestamp
				  time            time,            
		          screening_type  screening_type 
				  )`
	query[8] = &screening


	for  i := 0; i < 9; i++ {
		fmt.Printf("Executing query number %d = %s\n", i, *query[i] )
		_, err := db.Exec(*query[i])
		if err != nil {
			panic(err)
		}
		fmt.Println("Success!" )
	}
}


func RunProvisioning(){
	provisioningQueryMovies :=	`
							 INSERT INTO movie (movie_title,movie_year,pg_type, runtime)
							 VALUES 
							 ('Spider-Man: Far From Home', 2019,'G',138),
							 ('Fast & Furious: Hobbs & Shaw', 2019,'12',134),
							 ('The Lion King', 2019,'G',138),
							 ('Toy Story 4', 2019,'G',138),
							 ('Once Upon A Time In Hollywood', 2019,'18',161),
							 ('Scary Stories To Tell In The Dark: Unlimited Screening', 2019,'15',138),
							 ('Anabel Comes Home', 2019,'15',106);
							`

	_, err := db.Exec(provisioningQueryMovies)

	if err != nil {
		fmt.Println(fmt.Errorf("error during provisioning of the movies table", err))
	}

	provisioningQueryTheater :=	`
							 INSERT INTO theater ("name", "rows", floor,capacity)
							 VALUES 
							 ('The Pleasure Room', 4,1,100),
							 ('Think Twice Before Entering', 6,1,120),
							 ('Room 66', 6,2,60),
							 ('Basic', 6,10,160)
							`

	_, err = db.Exec(provisioningQueryTheater)

	if err != nil {
		fmt.Println(fmt.Errorf("error during provisioning the theather table", err))
	}




	provisioningQueryScreening :=	`
							 INSERT INTO screening (id_movie, id_theater,date,"time",screening_type)
							 VALUES 
							 (1, 2,'10-10-2020', '15:30','2D'),
							 (2, 3,'10-10-2020', '15:30','3D')
						
							`

	_, err = db.Exec(provisioningQueryScreening)

	if err != nil {
		fmt.Println(fmt.Errorf("error during provisioning the schedule table", err))
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