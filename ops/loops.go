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
		fmt.Printf("Event ID: %s /n", e.ID)

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
