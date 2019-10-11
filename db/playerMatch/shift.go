package playerMatch

import (
	"github.com/jimmson/teamteam/db/events"
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=create -updaters=round -table=player_match

var fsm = shift.NewFSM(events.GetTable()).
	Insert(MatchStatusMatchStarted, 	create{}, MatchStatusRoundStarted).

	Update(MatchStatusRoundStarted, 	round{}, MatchStatusRoundJoining, MatchStatusRoundFailed).
	Update(MatchStatusRoundJoining, 	round{}, MatchStatusRoundJoined, MatchStatusRoundFailed, MatchStatusRoundExcluded).
	Update(MatchStatusRoundJoined, 		round{}, MatchStatusRoundCollecting, MatchStatusRoundFailed).
	Update(MatchStatusRoundCollecting, 	round{}, MatchStatusRoundCollected, MatchStatusRoundFailed).
	Update(MatchStatusRoundCollected, 	round{}, MatchStatusRoundSubmitting, MatchStatusRoundFailed).
	Update(MatchStatusRoundSubmitting, 	round{}, MatchStatusRoundSubmitted, MatchStatusRoundFailed).
	Update(MatchStatusRoundSubmitted, 	round{}, MatchStatusRoundSuccess, MatchStatusRoundFailed).
	Update(MatchStatusRoundSuccess, 	round{}, MatchStatusRoundStarted, MatchStatusMatchEnded).

	Update(MatchStatusRoundFailed, 		round{}, MatchStatusRoundStarted, MatchStatusMatchEnded).
	Update(MatchStatusRoundExcluded, 	round{}, MatchStatusRoundStarted, MatchStatusMatchEnded).

	Update(MatchStatusMatchEnded, 		round{}).
	Build()

type create struct {
	Match 		int
	PlayerName	string
}

type round struct {
	ID 			int
	Rank 		int
	MyPart  	int
	PlayerPart 	int
}
