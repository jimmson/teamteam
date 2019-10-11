package ops

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/jimmson/teamteam/db/cursors"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
)

var (
	team    = flag.String("team", "teamteams", "team name")
	player  = flag.String("player", "teamteam", "player name")
	players = flag.Int("players", 4, "number of players in the team")
)

func StartLoops(b Backends) {
	go startMatchForever(b)
	go consumeEngine(b)
}

func startMatchForever(b Backends) {
	for {
		ctx := unsure.ContextWithFate(context.Background(), unsure.DefaultFateP())

		err := b.EngineClient().StartMatch(ctx, *team, *players)

		if errors.Is(err, engine.ErrActiveMatch) {
			// Match active, just ignore
		} else if err != nil {
			log.Error(ctx, errors.Wrap(err, "start match error"))
		} else {
			log.Info(ctx, "match started")
		}

		time.Sleep(time.Second * 10)
	}
}

func consumeEngine(b Backends) {

	f := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {
		fmt.Printf("Event: %s \n", String(e.Type.ReflexType()))

		return nil
	}

	// Internal engine events consumable.
	consumable := reflex.NewConsumable(
		b.EngineClient().Stream,
		cursors.ToStore(b.TeamteamDB().DB))

	//TODO: Better consumer name
	consumer := reflex.NewConsumer("player", f, []reflex.ConsumerOption{}...)

	unsure.ConsumeForever(unsureCtx, consumable.Consume, consumer, []reflex.StreamOption{}...)
}

func unsureCtx() context.Context {
	max := rand.Intn(60) // Max 60 secs.
	d := time.Second * time.Duration(max)
	ctx, cancel := context.WithTimeout(context.Background(), d)

	// Call cancel to satisfy golint.
	go func() {
		time.Sleep(d)
		cancel()
	}()

	return unsure.ContextWithFate(ctx, unsure.DefaultFateP())
}

// String returns the string representation of "t".
func String(t int) string {
	switch t {
	case engine.EventTypeUnknown.ReflexType():
		return "Unknown"
	case engine.EventTypeMatchStarted.ReflexType():
		return "Match Started"
	case engine.EventTypeMatchEnded.ReflexType():
		return "Match Ended"
	case engine.EventTypeRoundJoin.ReflexType():
		return "Round Joining"
	case engine.EventTypeRoundJoined.ReflexType():
		return "Round Joined"
	case engine.EventTypeRoundCollect.ReflexType():
		return "Round Collecting"
	case engine.EventTypeRoundCollected.ReflexType():
		return "Round Collected"
	case engine.EventTypeRoundSubmit.ReflexType():
		return "Round Submitting"
	case engine.EventTypeRoundSubmitted.ReflexType():
		return "Round Submitted"
	case engine.EventTypeRoundSuccess.ReflexType():
		return "Round Success"
	case engine.EventTypeRoundFailed.ReflexType():
		return "Round Failed"
	default:
		return "Invalid event type"
	}
}
