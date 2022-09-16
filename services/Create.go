package crud

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Team

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t.Id = uuid.New()
		s.FootbalLeague = append(s.FootbalLeague, t)
	}
}

// delete function
func (s *Server) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Team

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, v := range s.FootbalLeague {
			if v.Id == t.Id {
				s.FootbalLeague = append(s.FootbalLeague[:i], s.FootbalLeague[i+1:]...)
				return
			}
		}
	}
}

// update function
func (s *Server) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Team

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, v := range s.FootbalLeague {
			if v.Id == t.Id {
				s.FootbalLeague[i] = t
				return
			}
		}
	}
}
