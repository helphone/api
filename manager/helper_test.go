package manager

import (
	"strings"
	"testing"
)

func TestIsCountryExist(t *testing.T) {
	existingOne := "US"
	notExistingOne := "ZZ"

	if isCountryExist(existingOne) != true {
		t.Error("'US' country should exist in the database")
	}

	if isCountryExist(strings.ToLower(existingOne)) != false {
		t.Error("'us' country in lowercase shouldn't exist in the database")
	}

	if isCountryExist(notExistingOne) != false {
		t.Error("'ZZ' country shouldn't exist in the database")
	}
}

func TestIsLanguageExist(t *testing.T) {
	existingOne := "fr"
	notExistingOne := "zz"

	if isLanguageExist(existingOne) != true {
		t.Error("'fr' language should exist in the database")
	}

	if isLanguageExist(strings.ToUpper(existingOne)) != false {
		t.Error("'fr' language in upercase shouldn't exist in the database")
	}

	if isLanguageExist(notExistingOne) != false {
		t.Error("'zz' language shouldn't exist in the database")
	}
}
