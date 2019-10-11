package teamteam

import (
	"context"

	"github.com/luno/reflex"
)

// Client defines the root engine service interface.
type Client interface {
	Ping(context.Context) error
	Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error)
	GetRound(ctx context.Context, roundID int, playerName string) (*Round, error)
}
