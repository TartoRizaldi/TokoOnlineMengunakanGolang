package app

import "github.com/TartoRizaldi/TokoOnlineMengunakanGolang/app/controllers"

func (server *server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
