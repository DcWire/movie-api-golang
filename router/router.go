package router

import (
	"fmt"
	"net/http"

	"github.com/DcWire/mymodules/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Welcome /api/")
	}).Methods("GET")
	return router
}
