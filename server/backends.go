package server

import "github.com/jimmson/teamteam/db"

//go:generate genbackendsimpl
type Backends interface {
	TeamteamDB() *db.TeamteamDB
}
