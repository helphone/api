package manager

import (
	log "github.com/Sirupsen/logrus"
	"github.com/helphone/api/model"
)

// GetCountry will return a row for a specific country (used to know where you are)
func GetCountry(countryCode string, languageCode string) (*model.Response, error) {
	queryStmt, err := dB.Prepare("SELECT * FROM countries_by_language($1) WHERE country_code=$2")
	if err != nil {
		log.Errorf("Could not prepare the query to get countries, err: %s", err)
		return nil, err
	}

	rows, err := queryStmt.Query(languageCode, countryCode)
	defer rows.Close()

	if err != nil {
		log.Errorf("Error in the execution of the query, err: %s", err)
		return nil, err
	}

	return populateDataForCountries(rows)
}

// GetCountries will returns all countries founded
func GetCountries(languageCode string) (*model.Response, error) {
	queryStmt, err := dB.Prepare("SELECT * FROM countries_by_language($1)")
	if err != nil {
		log.Errorf("Could not prepare the query to get countries, err: %s", err)
		return nil, err
	}

	rows, err := queryStmt.Query(languageCode)
	defer rows.Close()

	if err != nil {
		log.Errorf("Error in the execution of the query, err: %s", err)
		return nil, err
	}

	return populateDataForCountries(rows)
}

// GetPhonenumbers will returns every phonenumbers founded with
// specific translations from the languageCode
func GetPhonenumbers(languageCode string) (*model.Response, error) {
	if languageCode != "" && isLanguageExist(languageCode) == false {
		return nil, ErrLanguageDoesNotExist
	}

	if languageCode == "" {
		languageCode = "en"
	}

	queryStmt, err := dB.Prepare("SELECT * FROM numbers_by_language($1)")
	if err != nil {
		log.Errorf("Could not prepare the query to get all phonenumbers, err: %s", err)
		return nil, err
	}

	rows, err := queryStmt.Query(languageCode)
	defer rows.Close()

	if err != nil {
		log.Errorf("Error in the execution of the query, err: %s", err)
		return nil, err
	}

	return populateDataForPhonenumbers(rows, languageCode)
}

// GetPhonenumbersForCountry will returns every phonenumbers founded
// for a country code with specific translations from the languageCode
func GetPhonenumbersForCountry(languageCode string, countryCode string) (*model.Response, error) {
	if countryCode != "" && isCountryExist(countryCode) == false {
		return nil, ErrLanguageDoesNotExist
	}

	if languageCode != "" && isLanguageExist(languageCode) == false {
		return nil, ErrLanguageDoesNotExist
	}

	if languageCode == "" {
		languageCode = "en"
	}

	queryStmt, err := dB.Prepare("SELECT * FROM numbers_by_language($2) WHERE country_code=$1")
	if err != nil {
		log.Errorf("Could not prepare the query to get all phonenumbers, err: %s", err)
		return nil, err
	}

	rows, err := queryStmt.Query(countryCode, languageCode)
	defer rows.Close()

	if err != nil {
		log.Errorf("Error in the execution of the query, err: %s", err)
		return nil, err
	}

	return populateDataForPhonenumbers(rows, languageCode)
}
