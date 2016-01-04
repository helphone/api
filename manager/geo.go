package manager

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/helphone/geo/georeverse"
)

var (
	geo *georeverse.CountryReverser
)

func prepareGeoReverse() {
	_geo, err := georeverse.NewCountryReverser("/etc/files/polygons.properties")
	if err != nil {
		log.Errorf("Could not prepare the geo reverse")
	}

	geo = _geo
}

// FindCountryFromLatAndLong allow you to know the contry code of a given
// latitude and longitude. If doesn't find any, the error will be set
func FindCountryFromLatAndLong(lat float64, long float64) (code string, err error) {
	code = strings.ToUpper(geo.GetCountryCode(long, lat))

	if code == "" {
		err = errors.New("No country found")
	}
	return
}
