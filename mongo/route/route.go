package route

import (
	"github.com/gorilla/mux"
	"github.com/vivekrqwat/monogo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movie", controller.Getall).Methods("GET")
	router.HandleFunc("/movie/{id}", controller.GetByid).Methods("GET")
	router.HandleFunc("/movie", controller.Create).Methods("POST")
	router.HandleFunc("/movie", controller.MarkasWatched).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleByid).Methods("DELETE")
	router.HandleFunc("/movie", controller.DeleAll).Methods("DELETE")

	return router
}
