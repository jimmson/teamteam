package cursors

import (
	"testing"

	"github.com/corverroos/unsure"
	"github.com/jimmson/teamteam/db"
	"github.com/luno/reflex/rsql"
)

func TestCursorsTable(t *testing.T) {
	defer unsure.CheatFateForTesting(t)()
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	rsql.TestCursorsTable(t, dbc, cursors)
}
