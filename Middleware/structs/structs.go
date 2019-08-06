package structs

// Bindings from and to JSON
type User struct {
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
}

// Binding struct
type Theater struct {
	Name    string
	Rows    string
	Floor   string
}

type NewMovie struct {
	Name      string
	MovieYear string
	PgType    string
	Runtime   string
	Theater   string
	Schedule  map[string][]string
}


