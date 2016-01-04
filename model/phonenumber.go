package model

// Phonenumber is the struct that define the return of the database
// and the formated JSON output
type Phonenumber struct {
	Category    string `json:"category" sql:"category_name"`
	Phonenumber string `json:"number" sql:"phone_number"`
	Country     string `json:"-" sql:"country_code"`
}
