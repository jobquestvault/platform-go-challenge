package sys

import "context"

type (
	Lifecycle interface {
		Setup(ctx context.Context) error
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)
