package model

// Response is the struct that define the return of the database
// and the formated JSON output
type Response struct {
	Countries []*Country `json:"countries"`
	Language  string     `json:"language,omitempty"`
}
