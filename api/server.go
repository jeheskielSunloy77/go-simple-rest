package api

import (
	"github.com/gorilla/mux"
	"github.com/jeheskielSunloy77/simpleserver/models"
)

func NewServer() *models.Server {
	s := &models.Server{
		Router:        mux.NewRouter(),
		FootbalLeague: models.Teams,
	}
	s.routes()
	return s
}

func (s *models.Server) routes() {
	s.HandleFunc("/teams", s.Create()).Methods("POST")
	s.HandleFunc("/teams", s.GetAll()).Methods("GET")
	s.HandleFunc("/teams/{id}", s.GetOne()).Methods("GET")
	s.HandleFunc("/teams/{id}", s.Update()).Methods("PUT")
	s.HandleFunc("/teams/{id}", s.Delete()).Methods("DELETE")
}
