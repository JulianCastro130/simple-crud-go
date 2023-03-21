package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"example/hello/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/names", handlers.GetNames).Methods("GET")
	r.HandleFunc("/names/{id}", handlers.GetName).Methods("GET")
	r.HandleFunc("/names", handlers.CreateName).Methods("POST")
	r.HandleFunc("/names/{id}", handlers.UpdateName).Methods("PUT")
	r.HandleFunc("/names/{id}", handlers.DeleteName).Methods("DELETE")

	return r
}

func StartServer(port string) error {
	r := NewRouter()

	return http.ListenAndServe(":"+port, r)
}
