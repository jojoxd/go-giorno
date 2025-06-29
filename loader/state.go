package loader

// State signifies a Controller state
type State[TData any] interface{}

// StateInitial is the initial State a Controller is in
type StateInitial[TData any] struct {
	State[TData]
}

// StateError is the State that a Controller contains when an error has occurred
type StateError[TData any] struct {
	State[TData]
	Error error
}

// StateQueued is the State that a Controller contains when a Controller.Load has been issued,
// but the async.Scheduler has not yet taken up the job yet
type StateQueued[TData any] struct {
	State[TData]
}

// StateLoading is the State that a Controller contains when the async.Scheduler has started loading
type StateLoading[TData any] struct {
	State[TData]
}

// StateLoaded is the State that a Controller contains when the async.Scheduler has finished loading
type StateLoaded[TData any] struct {
	State[TData]
	Data TData
}
