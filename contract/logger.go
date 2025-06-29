package contract

import (
	"log/slog"
)

// Expect that slog fulfills contract
var _ Logger = (*slog.Logger)(nil)

// Logger specifies a logging contract
type Logger interface {
	Warn(message string, args ...any)
	Info(message string, args ...any)
	Debug(message string, args ...any)
}
