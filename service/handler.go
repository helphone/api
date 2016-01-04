package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/helphone/api/manager"
	"github.com/helphone/api/model"
)

// JSONError is a simple struct that is used to render properly a message in JSON
type JSONError struct {
	Message string `json:"message"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request ping")

	fmt.Fprintf(w, "%s", "Pong")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	languageCode := query.Get("language")
	if languageCode == "" {
		languageCode = "en"
	}

	countryCode, err := getCountryFromLatAndLong(w, r)
	if err != nil {
		return
	}

	var data *model.Response

	if countryCode == "" {
		data, err = manager.GetCountries(languageCode)
	} else {
		data, err = manager.GetCountry(countryCode, languageCode)
	}

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	renderWithETag(w, r, response)
}

func getPhonenumbers(w http.ResponseWriter, r *http.Request) {
	var err error
	query := r.URL.Query()
	languageCode := query.Get("language")
	countryCode := query.Get("country")

	if countryCode == "" {
		countryCode, err = getCountryFromLatAndLong(w, r)
		if err != nil {
			return
		}
	} else {
		countryCode = strings.ToUpper(countryCode)
	}

	var data *model.Response
	if countryCode == "" {
		data, err = manager.GetPhonenumbers(languageCode)
	} else {
		data, err = manager.GetPhonenumbersForCountry(languageCode, countryCode)
		if len(data.Countries) == 0 {
			printANotFoundCountry(w, r)
			return
		}
	}

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	renderWithETag(w, r, response)
}
