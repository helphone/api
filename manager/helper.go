package manager

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	"github.com/helphone/api/model"
	"github.com/kisielk/sqlstruct"
)

func isCountryExist(code string) bool {
	err := dB.QueryRow("SELECT code FROM countries WHERE code = $1", code).Scan(new(string))
	return err == nil
}

func isLanguageExist(code string) bool {
	err := dB.QueryRow("SELECT code FROM languages WHERE code = $1", code).Scan(new(string))
	return err == nil
}

func populateDataForCountries(rows *sql.Rows) (*model.Response, error) {
	countries := []*model.Country{}

	for rows.Next() {
		var country model.Country
		err := sqlstruct.Scan(&country, rows)

		if err != nil {
			log.Errorf("Error in the parsing of a country, err: %s", err)
		}
		countries = append(countries, &country)
	}

	err := rows.Err()
	if err != nil {
		log.Errorf("Error during the loop of rows, err: %s", err)
	}

	data := &model.Response{
		Countries: countries,
		Language:  "",
	}
	return data, err
}

func populateDataForPhonenumbers(rows *sql.Rows, languageCode string) (*model.Response, error) {
	countries := []*model.Country{}

	for rows.Next() {
		var data model.Phonenumber
		err := sqlstruct.Scan(&data, rows)

		if err != nil {
			panic(err)
		}

		countryPosition := -1
		for i := 0; i < len(countries); i++ {
			if countries[i].Code == data.Country {
				countryPosition = i
			}
		}

		if countryPosition != -1 {
			countries[countryPosition].Phonenumbers = append(countries[countryPosition].Phonenumbers, &data)
		} else {
			newCountry := &model.Country{
				Code:         data.Country,
				Phonenumbers: []*model.Phonenumber{},
			}
			newCountry.Phonenumbers = append(newCountry.Phonenumbers, &data)
			countries = append(countries, newCountry)
		}
	}

	err := rows.Err()
	if err != nil {
		log.Errorf("Error during the loop of rows, err: %s", err)
	}

	data := &model.Response{
		Countries: countries,
		Language:  languageCode,
	}
	return data, err
}
