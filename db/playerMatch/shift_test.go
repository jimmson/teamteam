package playerMatch

import (
	"github.com/jimmson/teamteam/db"
	"testing"

	"github.com/luno/shift"
	"github.com/stretchr/testify/require"
	)

func TestFSM(t *testing.T) {
	dbc := db.ConnectForTesting(t)
	defer dbc.Close()

	require.NoError(t, shift.TestFSM(t, dbc, fsm))
}
