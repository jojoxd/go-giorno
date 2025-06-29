package router

type BaseScreen struct {
	Intent Intent
}

func (s *BaseScreen) Location() RouteLocation {
	return s.Intent.Location()
}

func (s *BaseScreen) OnIntent(intent Intent) error {
	s.Intent = intent

	return nil
}
