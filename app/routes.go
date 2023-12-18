package app

import (
	"pasar-gelap/app/controllers"
	"pasar-gelap/app/middlewares/log"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()

	// logger middleware
	server.Router.Use(log.LoggerMiddleware)

	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
