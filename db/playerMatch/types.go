package playerMatch

import "time"

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
	ID        int64               `protocp:"1"`
	Status    MatchStatus         `protocp:"2"`
	CreatedAt time.Time           `protocp:"3"`
	UpdatedAt time.Time           `protocp:"4"`
}

type MatchJSON struct {
	ID        int64         `json:"id"`
	Status    MatchStatus 	`json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
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

var matchReflex = map[MatchStatus]EventType{
	// TODO(Kyle): Add Reflex Event types
}

