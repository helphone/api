package model

// Country is the struct that define the return of the database
// and the formated JSON output
type Country struct {
	Code         string         `json:"code" sql:"country_code"`
	Name         string         `json:"name,omitempty" sql:"country_name"`
	Phonenumbers []*Phonenumber `json:"phonenumbers,omitempty"`
}
