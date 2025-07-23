package handlers

import (
	"net/http"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// Service groups customer methods.
type Service struct{}

var customers = map[int]Customer{}

func (s *Service) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Service) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (s *Service) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
