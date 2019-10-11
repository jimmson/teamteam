package playerMatch

import (
	"context"
	"database/sql"
)

const cols = " id, status, created_at, updated_at, player_name, round_num, rank, my_part, player_part "

// Create starts the match
func Create(ctx context.Context, dbc *sql.DB, match int) (int64, error) {
	return fsm.Insert(ctx, dbc, create{Match: match})
}

// -----------------
// Try Starting Round
func StartRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusMatchStarted, MatchStatusRoundStarted,
		roundMove{ID: id})
}

// -----------------
// Failed Round
func RoundFailed(ctx context.Context, dbc *sql.DB, id int64, status MatchStatus) error {
	return fsm.Update(ctx, dbc, status, MatchStatusRoundFailed,
		roundMove{ID: id})
}

// -----------------
// End Match
func EndMatch(ctx context.Context, dbc *sql.DB, id int64, status MatchStatus) error {
	return fsm.Update(ctx, dbc, status, MatchStatusMatchEnded,
		roundMove{ID: id})
}

// -----------
// Try Joining Round
func JoiningRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundStarted, MatchStatusRoundJoining,
		roundMove{ID: id})
}

// -----------
// Try confirm Joined Round
func JoinedRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundJoining, MatchStatusRoundJoined,
		roundMove{ID: id})
}

func ExcludedFromRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundJoining, MatchStatusRoundExcluded,
		roundMove{ID: id})
}

// -----------
// Try Collecting Round
func CollectingRound(ctx context.Context, dbc *sql.DB, id int64, roundNum, rank, myparts, playerpart int) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundJoined, MatchStatusRoundCollecting,
		roundCollect{ID: id, RoundNum: roundNum, Rank: rank, MyPart: myparts, PlayerPart: playerpart})
}

// -----------
// Try Confirm Collected Round
func CollectedRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundCollecting, MatchStatusRoundCollected,
		roundMove{ID: id})
}

// ----------------
// Try Submitting Round
func SubmittingRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundCollected, MatchStatusRoundSubmitting,
		roundMove{ID: id})
}

// -----------
// Try confirm Submitted Round
func SubmittedRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundSubmitting, MatchStatusRoundSubmitted,
		roundMove{ID: id})
}

// -----------
// Try Succeed Round
func SuccessRound(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundSubmitted, MatchStatusRoundSuccess,
		roundMove{ID: id})
}

// -----------
// Succeed Round
func SucceedRoundButNewMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundSuccess, MatchStatusRoundStarted,
		roundMove{ID: id})
}

func SucceedRoundAndEndMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundFailed, MatchStatusMatchEnded,
		roundMove{ID: id})
}

// -----------
// Fail Round
func FailRoundButNewMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundFailed, MatchStatusRoundStarted,
		roundMove{ID: id})
}

func FailRoundAndEndMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundFailed, MatchStatusMatchEnded,
		roundMove{ID: id})
}

// -----------
// Excluded from Round
func ExcludeRoundButNewMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundExcluded, MatchStatusRoundStarted,
		roundMove{ID: id})
}

func ExcludeRoundAndEndMatch(ctx context.Context, dbc *sql.DB, id int64) error {
	return fsm.Update(ctx, dbc, MatchStatusRoundExcluded, MatchStatusMatchEnded,
		roundMove{ID: id})
}

func Lookup(ctx context.Context, dbc *sql.DB, id int64) (*Match, error) {
	return scan(dbc.QueryRowContext(ctx, "select "+cols+" from player_match where id=?", id))
}

func scan(row *sql.Row) (*Match, error) {
	var r Match

	err := row.Scan(&r.ID, &r.Status, &r.CreatedAt, &r.UpdatedAt, &r.PlayerName, &r.RoundNum, &r.Rank, &r.MyPart, &r.PlayerPart)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
