package intent

import (
	"git.jojoxd.nl/projects/go-giorno/router/internal"
)

type Base interface {
	Target() internal.Target
	implementsIntent()
}
