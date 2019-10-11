package ops

import (
	"github.com/corverroos/unsure/engine"
	"github.com/jimmson/teamteam/db"
)

//go:generate genbackendsimpl
type Backends interface {
	EngineClient() engine.Client
	TeamteamDB() *db.TeamteamDB
}
