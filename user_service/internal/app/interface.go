package app

import "context"

type App interface {
	Start(ctx context.Context) error
	Shutdown() error
}
