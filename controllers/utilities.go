package controllers

import (
	"database/sql"
)

type SQLPhonenumber struct {
	Category    string `json:"category" sql:"category_name"`
	Phonenumber string `json:"phonenumber" sql:"phone_number"`
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
