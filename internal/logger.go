package internal

import (
	"fmt"

	"git.jojoxd.nl/projects/go-giorno/contract"
)

// nilLogger is a logger that will not log anything (like /dev/null)
type nilLogger struct{}

func (nilLogger) Warn(_ string, _ ...any)  {}
func (nilLogger) Info(_ string, _ ...any)  {}
func (nilLogger) Debug(_ string, _ ...any) {}

func NewNilLogger() contract.Logger {
	return nilLogger{}
}

// prefixLogger is a logger that will prefix a string onto a log message
type prefixLogger struct {
	prefix string
	inner  contract.Logger
}

func NewPrefixLogger(prefix string, inner contract.Logger) contract.Logger {
	return prefixLogger{
		prefix: prefix,
		inner:  inner,
	}
}

func (p prefixLogger) Warn(message string, args ...any) {
	p.inner.Warn(fmt.Sprintf("%s: %s", p.prefix, message), args...)
}

func (p prefixLogger) Info(message string, args ...any) {
	p.inner.Info(fmt.Sprintf("%s: %s", p.prefix, message), args...)
}

func (p prefixLogger) Debug(message string, args ...any) {
	p.inner.Debug(fmt.Sprintf("%s: %s", p.prefix, message), args...)
}
