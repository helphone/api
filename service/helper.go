package service

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/helphone/api/manager"
)

const (
	countryNotFoundMessage            = "This country doesn't exist"
	missingCountryOrLatAndLongMessage = "You must provide a country or lat and long coordinates"
	latAndLongNeedToBeNumber          = "lat and long parameters need to be numbers"
)

func renderWithETag(w http.ResponseWriter, r *http.Request, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method == "GET" && len(data) > 1024 {
		hash := md5.New()
		for header := range w.Header() {
			io.WriteString(hash, header)
		}
		io.WriteString(hash, string(data))
		etag := fmt.Sprintf("%x", hash.Sum(nil))

		if match, ok := r.Header["If-None-Match"]; ok {
			if match[0] == etag {
				w.WriteHeader(http.StatusNotModified)
				data = []byte{}
			}
		} else {
			w.Header().Set("ETag", etag)
		}
	}

	w.Write(data)
}

func getCountryFromLatAndLong(w http.ResponseWriter, r *http.Request) (countryCode string, err error) {
	query := r.URL.Query()
	latString := query.Get("lat")
	longString := query.Get("long")

	if (latString != "" && longString == "") || (latString == "" && longString != "") {
		printMissingLatAndLong(w, r)
		return countryCode, errors.New(missingCountryOrLatAndLongMessage)
	} else if latString != "" && longString != "" {
		lat, err := strconv.ParseFloat(latString, 64)
		if err != nil {
			printLatAndLongNeedToBeNumber(w, r)
			return countryCode, errors.New(latAndLongNeedToBeNumber)
		}

		long, err := strconv.ParseFloat(longString, 64)
		if err != nil {
			printLatAndLongNeedToBeNumber(w, r)
			return countryCode, errors.New(latAndLongNeedToBeNumber)
		}

		countryCode, err = manager.FindCountryFromLatAndLong(lat, long)
		if err != nil {
			printANotFoundCountry(w, r)
			return countryCode, errors.New(countryNotFoundMessage)
		}
	}

	return countryCode, nil
}

func printANotFoundCountry(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	data, _ := json.Marshal(JSONError{countryNotFoundMessage})
	w.Write(data)
	return
}

func printLatAndLongNeedToBeNumber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	response, _ := json.Marshal(JSONError{latAndLongNeedToBeNumber})
	fmt.Fprintf(w, "%s", response)
	return
}

func printMissingLatAndLong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	response, _ := json.Marshal(JSONError{missingCountryOrLatAndLongMessage})
	fmt.Fprintf(w, "%s", response)
	return
}
