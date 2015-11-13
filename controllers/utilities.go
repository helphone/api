package controllers

import (
	"database/sql"
)

type Phonenumber struct {
	Category    string `json:"category" sql:"category_name"`
	Phonenumber string `json:"number" sql:"phone_number"`
	Country     string `json:"-" sql:"country_code"`
}

type Country struct {
	Code         string         `json:"code"`
	Phonenumbers []*Phonenumber `json:"phonenumbers"`
}

type JSONResponse struct {
	Language  string     `json:"language"`
	Countries []*Country `json:"countries"`
}

func isCountryExist(code *string, db *sql.DB) bool {
	err := db.QueryRow("SELECT code FROM countries WHERE code = $1", code).Scan(new(string))

	if err != nil {
		return false
	} else {
		return true
	}
}

func isLanguageExist(code *string, db *sql.DB) bool {
	err := db.QueryRow("SELECT code FROM languages WHERE code = $1", code).Scan(new(string))

	if err != nil {
		return false
	} else {
		return true
	}
}
