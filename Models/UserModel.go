package Models

//this is user struct
type User struct {
	FirstName string        `json:"firstname"`
	LastName string         `json:"lastname"`
	Address *Address `json:"address"`
}

