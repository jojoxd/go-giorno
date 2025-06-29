package main

import (
	"gioui.org/layout"

	"git.jojoxd.nl/projects/go-giorno/example/base"
)

func main() {
	ex := example{}
	base.Run(ex.frame)
}

type example struct{}

func (e example) frame(gtx layout.Context) {

}
