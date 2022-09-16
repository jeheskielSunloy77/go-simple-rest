package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		for i, team := range s.FootbalLeague {
			if team.Id == id {
				s.FootbalLeague = append(s.FootbalLeague[:i], s.FootbalLeague[i+1:]...)
				break
			}
		}
	}
}
