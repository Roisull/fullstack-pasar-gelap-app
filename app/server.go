package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) initialize() {
	fmt.Println("Selamat datang di Pasar Gelap")

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string){
	fmt.Printf("Running on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run(){
	var server = Server{}
	server.initialize()
	server.Run(":9000")
}