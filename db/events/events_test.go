package events

import (
	"testing"

	"github.com/corverroos/unsure"
	"github.com/jimmson/teamteam/db"
	"github.com/luno/reflex/rsql"
)

func TestEventsTable(t *testing.T) {
	defer unsure.CheatFateForTesting(t)()
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	rsql.TestEventsTableInt(t, dbc, events)
}
