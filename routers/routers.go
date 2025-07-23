package routers

import (
	"github.com/gorilla/mux"

	"api/handlers"
)

func NewRouter() *mux.Router {
	routes := mux.NewRouter()

	svc := &handlers.Service{}

	routes.HandleFunc("/customers", svc.GetCustomers).Methods("GET")
	routes.HandleFunc("/customers/{id}", svc.GetCustomer).Methods("GET")
	routes.HandleFunc("/customers", svc.CreateCustomer).Methods("POST")

	return routes
}
