package playerMatch

import (
	"github.com/corverroos/unsure/engine"
	"time"
)

type MatchStatus int

func (ms MatchStatus) Enum() int {
	return int(ms)
}

func (MatchStatus) ShiftStatus() {
}

func (ms MatchStatus) ReflexType() int {
	return matchReflex[ms].ReflexType()
}

type Match struct {
	ID        		int64
	Status    		MatchStatus
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	MatchID 	  	int
	PlayerName		string
	RoundNum		int
	Rank 			int
	MyPart  		int
	PlayerPart 		int
}

type MatchJSON struct {
	ID        int64         `json:"id"`
	Status    MatchStatus 	`json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	MatchID 	  	int		`json:"match_id"`
	PlayerName		string	`json:"player_Name"`
	RoundNum		int		`json:"round_num"`
	Rank 			int		`json:"rank"`
	MyPart  		int		`json:"my_part"`
	PlayerPart 		int		`json:"player_part"`
}

var (
	MatchStatusUnknown  	 			MatchStatus = 0
	MatchStatusMatchStarted   			MatchStatus = 1

	//Round
	MatchStatusRoundStarted				MatchStatus = 2
	MatchStatusRoundJoining 			MatchStatus = 3
	MatchStatusRoundJoined	  			MatchStatus = 4
	MatchStatusRoundCollecting			MatchStatus = 5
	MatchStatusRoundCollected	  		MatchStatus = 6
	MatchStatusRoundSubmitting	  		MatchStatus = 7
	MatchStatusRoundSubmitted	  		MatchStatus = 8
	MatchStatusRoundSuccess			  	MatchStatus = 9
	//Failure and exclusion
	MatchStatusRoundFailed			  	MatchStatus = 10
	MatchStatusRoundExcluded	  		MatchStatus = 11
	//Match ended
	MatchStatusMatchEnded				MatchStatus = 12
)

var matchReflex = map[MatchStatus]engine.EventType{
	MatchStatusMatchStarted: engine.EventTypeMatchStarted,
	MatchStatusRoundStarted: engine.EventTypeRoundJoin,
	MatchStatusRoundJoined: engine.EventTypeRoundJoined,
	MatchStatusRoundCollecting: engine.EventTypeRoundCollect,
	MatchStatusRoundCollected: engine.EventTypeRoundCollected,
	MatchStatusRoundSubmitting:engine.EventTypeRoundSubmit,
	MatchStatusRoundSubmitted: engine.EventTypeRoundSubmitted,
	MatchStatusRoundSuccess: engine.EventTypeRoundSuccess,
	MatchStatusRoundFailed: engine.EventTypeRoundFailed,
	MatchStatusMatchEnded: engine.EventTypeMatchEnded,
}

