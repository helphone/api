package manager

import (
	"testing"
)

func init() {
	Init()
	<-ReadyChannel
}

func TestPingDatabase(t *testing.T) {
	err := dB.Ping()
	if err != nil {
		t.Error("Can't ping the database")
	}
}

func TestGetCountry(t *testing.T) {
	response, err := GetCountry("BE", "en")
	if err != nil {
		t.Errorf("Can't get country, err: %s", err)
	}

	if len(response.Countries) != 1 {
		t.Error("Countries list shouldn't be empty")
	}

	if response.Language != "" {
		t.Error("Language code should be an empty string")
	}
}

func TestGetCountries(t *testing.T) {
	response, err := GetCountries("en")
	if err != nil {
		t.Errorf("Can't get countries, err: %s", err)
	}

	if len(response.Countries) == 0 {
		t.Error("Countries list shouldn't be empty")
	}

	if response.Language != "" {
		t.Error("Language code should be an empty string")
	}
}

func TestGetPhonenumbers(t *testing.T) {
	response, err := GetPhonenumbers("en")
	if err != nil {
		t.Errorf("Can't get phonenumbers, err: %s", err)
	}

	if len(response.Countries) == 0 {
		t.Error("Countries list shouldn't be empty")
	}

	if response.Language != "en" {
		t.Error("Language code should be an empty string")
	}

	_, err = GetPhonenumbers("aze")
	if err == nil {
		t.Error("It should return an error for a non-present language")
	}
}

func TestGetPhonenumbersForCountry(t *testing.T) {
	response, err := GetPhonenumbersForCountry("en", "BE")
	if err != nil {
		t.Errorf("Can't get phonenumbers, err: %s", err)
	}

	if len(response.Countries) == 0 {
		t.Error("Countries list shouldn't be empty")
	}

	if response.Language != "en" {
		t.Error("Language code should be an empty string")
	}

	_, err = GetPhonenumbersForCountry("aze", "BE")
	if err == nil {
		t.Error("It should return an error for a non-present language")
	}

	_, err = GetPhonenumbersForCountry("en", "AZE")
	if err == nil {
		t.Error("It should return an error for a non-present country")
	}

	_, err = GetPhonenumbersForCountry("en", "be")
	if err == nil {
		t.Error("It should return an error for a non-uppercase country")
	}
}
