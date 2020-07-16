package inquiry

import (
	"github.com/gorilla/mux"
	inq "tapera.integrasi/service/inquiry"
)

// Controller struct
type Controller struct {
	srv inq.Service
}

// NewController func
func NewController(srv inq.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	//routes grouping
	s := r.PathPrefix("/dummy/bri").Subrouter()
	s.HandleFunc("/test", c.Test).Methods("GET")
	s.HandleFunc("/inquiry", c.Inquiry).Methods("POST")
}
