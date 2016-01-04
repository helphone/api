package manager

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	// Load postgres driver
	_ "github.com/lib/pq"
)

var (
	dB *sql.DB

	// ReadyChannel is a channel to know when the manager is ready
	ReadyChannel = make(chan int, 1)
	// ErrCountryDoesNotExist is a custom error
	ErrCountryDoesNotExist = errors.New("The country doesn't exist")
	// ErrLanguageDoesNotExist is a custom error
	ErrLanguageDoesNotExist = errors.New("The language doesn't exist")
)

// Init is a function to prepare the service
func Init() {
	prepareDatabase()
	prepareGeoReverse()

	ReadyChannel <- 1
}

func prepareDatabase() {
	hostname := os.Getenv("DB_PORT_5432_TCP_ADDR")
	port := os.Getenv("DB_PORT_5432_TCP_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, "helphone")

	_dB, err := sql.Open("postgres", dbInfo)
	if err != nil && _dB.Ping() != nil {
		log.Errorf("Could not connect to the database")
	}

	dB = _dB
}
