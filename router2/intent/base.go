package intent

import (
	"git.jojoxd.nl/projects/go-giorno/router2/internal"
)

type Base interface {
	Target() internal.Target
	implementsIntent()
}
