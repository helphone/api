package manager

import (
	"testing"
)

func TestFindCountryFromLatAndLong(t *testing.T) {
	res, err := FindCountryFromLatAndLong(100.0, 100.0)
	if res != "" || err == nil {
		t.Errorf("The country should not exist. Find %s", res)
	}

	res, err = FindCountryFromLatAndLong(50.625073, 4.746094)
	if res != "BE" || err != nil {
		t.Errorf("The country should exist and be 'BE'. Find %s", res)
	}
}
