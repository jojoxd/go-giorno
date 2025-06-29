package history

import (
	"git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

// History contains the contract to manage the router stack History
type History interface {
	// Push pushes an Item onto a History stack
	Push(*Item)

	// Pop pulls the current Item off of a History stack
	Pop() *Item

	// Peek returns the current of the Item of a History stack
	Peek() *Item

	// Depth returns the total depth of the History stack
	Depth() int

	// All returns all the Item's from the History stack
	All() []*Item

	// Clear clears all the Item's from the History stack
	Clear()
}

// Item stores a reference to a view.View based on an intent.Base
type Item struct {
	View   view.View
	Intent intent.Base
}
