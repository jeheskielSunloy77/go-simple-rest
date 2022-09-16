package models

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Teams = []Team{
	{
		Id:       uuid.New(),
		TeamName: "Real Madrid",
	},
	{
		Id:       uuid.New(),
		TeamName: "Barcelona",
	},
	{
		Id:       uuid.New(),
		TeamName: "Juventus",
	},
	{
		Id:       uuid.New(),
		TeamName: "Bayern Munich",
	},
	{
		Id:       uuid.New(),
		TeamName: "Liverpool FC",
	},
	{
		Id:       uuid.New(),
		TeamName: "AC Milan FC",
	},
}

type Team struct {
	Id       uuid.UUID `json:"id"`
	TeamName string    `json:"teamName"`
}
type Server struct {
	*mux.Router
	FootbalLeague []Team
}
