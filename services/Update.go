package crud

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jeheskielSunloy77/simpleserver/models"
)

func (s *models.Server) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Team
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, team := range s.FootbalLeague {
			if team.Id == id {
				t = Team{Id: id}
				if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				s.FootbalLeague[i] = t
				break
			}
		}

	}
}
