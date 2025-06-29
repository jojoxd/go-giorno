package fixedpool

import (
	"context"
	"fmt"
	"sync"

	"gioui.org/app"

	"git.jojoxd.nl/projects/go-giorno/async"
)

type scheduler struct {
	window *app.Window
	config *config
	workCh chan async.Schedulable
	init   sync.Once
}

func NewScheduler(window *app.Window, opts ...Option) async.Scheduler {
	conf := newDefaultConfig()
	conf.Load(opts...)

	return &scheduler{
		window: window,
		config: conf,
		workCh: make(chan async.Schedulable),
	}
}

func (s *scheduler) Schedule(ctx context.Context, fn async.Schedulable) {
	s.init.Do(func() {
		s.config.logger.Debug(fmt.Sprintf("starting %d workers", s.config.workers))

		for i := 0; i < s.config.workers; i++ {
			go func() {
				s.config.logger.Debug(fmt.Sprintf("starting worker %d", i))

				for work := range s.workCh {
					if work != nil {
						s.config.logger.Debug(fmt.Sprintf("worker %d: found some work", i))
						s.window.Invalidate()

						work.Execute(ctx)

						s.config.logger.Debug(fmt.Sprintf("worker %d: work complete", i))
						s.window.Invalidate()
					}
				}
			}()
		}
	})

	s.workCh <- fn
}
