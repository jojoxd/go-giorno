package view

type Finisher interface {
	Finish()
}

func Finish(view View) {
	if finisher, ok := view.(Finisher); ok {
		finisher.Finish()
	}
}
