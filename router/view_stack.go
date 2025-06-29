package router

import (
	"container/list"
	"iter"

	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
)

// ViewStack is for view navigation history
type ViewStack struct {
	viewList *list.List
	logger   contract.Logger
}

func (vs *ViewStack) Pop() RouteView {
	head := vs.viewList.Front()
	if head == nil {
		return nil
	}

	return vs.viewList.Remove(head).(RouteView)
}

func (vs *ViewStack) Peek() RouteView {
	if vs.viewList == nil || vs.viewList.Len() <= 0 {
		return nil
	}

	if vs.viewList.Front() == nil {
		return nil
	}

	return vs.viewList.Front().Value.(RouteView)
}

// push a new view to the stack and removes duplicates instance of the same view.
func (vs *ViewStack) Push(vw RouteView) error {
	if vs.viewList == nil {
		vs.viewList = list.New()
		vs.viewList.Init()
	}

	// v := vs.viewList.Front()
	// for v != nil {
	// 	if existing := v.Value.(RouteView); existing.Id() == vw.Id() {
	// 		vs.viewList.Remove(v)
	// 	}
	// 	v = v.Next()
	// }

	vs.viewList.PushFront(vw)
	return nil
}

func (vs *ViewStack) IsEmpty() bool {
	return vs.viewList == nil || vs.viewList.Len() <= 0
}

func (vs *ViewStack) Depth() int {
	if vs.viewList == nil {
		return 0
	}
	return vs.viewList.Len()
}

// All returns a iterator that iterates through the stack of routeProviders from back to front,
// or from front to back.
func (vs *ViewStack) All(backward bool) iter.Seq[RouteView] {
	return func(yield func(RouteView) bool) {
		if vs.viewList == nil || vs.viewList.Len() <= 0 {
			return
		}

		var v *list.Element
		if backward {
			v = vs.viewList.Back()
		} else {
			v = vs.viewList.Front()
		}
		for v != nil {
			if !yield(v.Value.(RouteView)) {
				return
			}

			if backward {
				v = v.Prev()
			} else {
				v = v.Next()
			}
		}
	}
}

func (vs *ViewStack) Clear() {
	if vs.viewList == nil {
		return
	}

	v := vs.viewList.Front()
	for v != nil {
		val := v.Value.(RouteView)

		finishRouteView(val, vs.logger)
		v = v.Next()
	}

	vs.viewList.Init()
}

func NewViewStack(logger contract.Logger) *ViewStack {
	if logger == nil {
		logger = internal.NewNilLogger()
	}

	return &ViewStack{
		viewList: &list.List{},
		logger:   logger,
	}
}
