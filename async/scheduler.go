package async

import (
	"context"
)

func (fn ScheduleFn) Execute(ctx context.Context) {
	fn(ctx)
}

type Scheduler interface {
	Schedule(context.Context, Schedulable)
}
