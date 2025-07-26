package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// groups customer logic methods.
type Service struct{}

// var customers = map[int]Customer{} // i tried to use map but 
// it was not the best way to do it so i used slice
var customers []Customer

func (s *Service) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, "Failed to encode customers", http.StatusInternalServerError)
		return
	}
}

func (s *Service) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	for _, customer := range customers {
		if customer.ID == id {
			if err := json.NewEncoder(w).Encode(customer); err != nil {
				http.Error(w, "Failed to encode customer", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "Customer not found", http.StatusNotFound)
}

func (s *Service) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer Customer
	json.NewDecoder(r.Body).Decode(&newCustomer) // Covert Json recieved from user to Go struct
	newCustomer.ID = int(uuid.New().ID())
	customers = append(customers, newCustomer)
	json.NewEncoder(w).Encode(newCustomer)
}
