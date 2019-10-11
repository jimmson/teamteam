package state

import (
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"
	"github.com/jimmson/teamteam/db"
)

type State struct {
	engineClient engine.Client
	teamteamDB   *db.TeamteamDB
}

func (s *State) EngineClient() engine.Client {
	return s.engineClient
}

func (s *State) TeamteamDB() *db.TeamteamDB {
	return s.teamteamDB
}

// New returns a new engine state.
func New() (*State, error) {
	var (
		s   State
		err error
	)

	s.engineClient, err = ec.New()
	if err != nil {
		return nil, err
	}

	s.teamteamDB, err = db.Connect()
	if err != nil {
		return nil, err
	}

	return &s, nil
}
