package service

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/helphone/api/manager"
)

// MuxWrapper is a struct that will wrap the router and the ready state
type MuxWrapper struct {
	IsReady bool
	Router  *mux.Router
}

func (httpWrapper *MuxWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	select {
	case <-manager.ReadyChannel:
		httpWrapper.IsReady = true
	default:
	}

	if httpWrapper.IsReady {
		httpWrapper.Router.ServeHTTP(w, r)
	} else {
		log.Debug("Service Unavailable")
		httpWrapper.returnCode503(w, r)
	}
}

func (httpWrapper *MuxWrapper) returnCode503(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("API Service is not yet available, please try again later"))
}

// NewRouter will return a new router with defined routes to functions
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ping)
	router.HandleFunc("/countries", getCountries)
	router.HandleFunc("/phonenumbers", getPhonenumbers)

	return router
}
