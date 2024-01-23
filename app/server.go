package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//server go di gunakan untuk booting sistem

type server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *server) Initialize() {
	fmt.Println("Welcome Toko online Rizal")

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *server) Run(addr string) {
	fmt.Printf("listening to port  %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
func Run() {
	var server = server{}

	server.Initialize()
	server.Run(":9000")
}
