package async

import "context"

type Schedulable interface {
	Execute(ctx context.Context)
}

type ScheduleFn func(ctx context.Context)
