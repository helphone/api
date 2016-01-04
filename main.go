package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/helphone/api/manager"
	"github.com/helphone/api/service"
	"github.com/rs/cors"
)

const (
	port = "3000"
)

func main() {
	log.Infof("Starting Helphone API service on port %s", port)

	router := service.NewRouter()
	h := service.MuxWrapper{
		IsReady: false,
		Router:  router,
	}

	go manager.Init()

	handler := handlers.CompressHandler(handlers.ProxyHeaders(cors.Default().Handler(h.Router)))
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
