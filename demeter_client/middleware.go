package demeter_client

import (
	"context"

	faktory "gitlab.id.vin/gami/demeter/client"
)

type Handler func(ctx context.Context, job *faktory.Job) error
type MiddlewareFunc func(ctx context.Context, job *faktory.Job, next func(ctx context.Context) error) error

// Use(...) adds middleware to the chain.
func (mgr *Manager) Use(middleware ...MiddlewareFunc) {
	mgr.middleware = append(mgr.middleware, middleware...)
}

func dispatch(chain []MiddlewareFunc, ctx context.Context, job *faktory.Job, perform Handler) error {
	if len(chain) == 0 {
		return perform(ctx, job)
	}

	link := chain[0]
	rest := chain[1:]
	return link(ctx, job, func(ctx context.Context) error {
		return dispatch(rest, ctx, job, perform)
	})
}
