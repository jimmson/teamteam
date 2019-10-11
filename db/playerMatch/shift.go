package playerMatch

import (
	"github.com/jimmson/teamteam/db/events"
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=create -updaters=roundMove,roundCollect -table=player_match

var fsm = shift.NewFSM(events.GetTable()).
	Insert(MatchStatusMatchStarted, 	create{}, MatchStatusRoundStarted).

	Update(MatchStatusRoundStarted, 	roundMove{}, MatchStatusRoundJoining, MatchStatusRoundFailed).
	Update(MatchStatusRoundJoining, 	roundMove{}, MatchStatusRoundJoined, MatchStatusRoundFailed, MatchStatusRoundExcluded).
	Update(MatchStatusRoundJoined, 		roundMove{}, MatchStatusRoundCollecting, MatchStatusRoundFailed).
	Update(MatchStatusRoundCollecting, 	roundCollect{}, MatchStatusRoundCollected, MatchStatusRoundFailed).
	Update(MatchStatusRoundCollected, 	roundMove{}, MatchStatusRoundSubmitting, MatchStatusRoundFailed).
	Update(MatchStatusRoundSubmitting, 	roundMove{}, MatchStatusRoundSubmitted, MatchStatusRoundFailed).
	Update(MatchStatusRoundSubmitted, 	roundMove{}, MatchStatusRoundSuccess, MatchStatusRoundFailed).
	Update(MatchStatusRoundSuccess, 	roundMove{}, MatchStatusRoundStarted, MatchStatusMatchEnded).

	Update(MatchStatusRoundFailed, 		roundMove{}, MatchStatusRoundStarted, MatchStatusMatchEnded).
	Update(MatchStatusRoundExcluded, 	roundMove{}, MatchStatusRoundStarted, MatchStatusMatchEnded).

	Update(MatchStatusMatchEnded, 		roundMove{}).
	Build()

type create struct {
	Match 		int
}

type roundCollect struct {
	ID 			int64
	RoundNum	int
	Rank 		int
	MyPart  	int
	PlayerPart 	int
}

type roundMove struct {
	ID 			int64
}
