package server

import (
	"OrderManagement/config"
	"OrderManagement/repository"
	"OrderManagement/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
}

type server struct {
}

func (sr server) Start() error {
	dbConfig := config.GetDBConfig()
	dbConn, err := repository.GetGORMDBConnection(dbConfig)
	if err != nil {
		return fmt.Errorf("GormClient: %s", err.Error())
	}
	// setup Router
	router := mux.NewRouter()

	// DB client setup
	dbClient := repository.NewGORMClient(dbConn)
	// Setup Repository
	ordersRespository := repository.NewOrdersRepository(dbClient)
	orderItemsRepository := repository.NewOrderItemRepository(dbClient)
	// Setup Services and Handlers
	routesService := routes.NewRoutesService(ordersRespository, orderItemsRepository)
	routes.NewRoutesHandler(router, routesService)

	fmt.Println("Listening on Port:", dbConfig.Port)
	log.Fatal(http.ListenAndServe(":"+dbConfig.Port, router))
	return nil

}

func NewServer() Server {
	return server{}
}
