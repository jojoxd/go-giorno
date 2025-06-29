package loader

type Controller[TArg comparable, TData any] interface {
	Load(TArg)
	State() State[TData]
	Data() (data TData, ok bool)
}
