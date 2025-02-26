package logger

import (
	"go.uber.org/fx/fxevent"
	"strings"
	"weather-cache/config"
)

// NewFxLogger returns a new fxevent.Logger that logs to the provided logger.
func NewFxLogger(l Logger, config *config.Config) fxevent.Logger {
	if strings.ToLower(config.Log.EnableFxLogs) == "yes" {
		return &fxevent.SlogLogger{
			Logger: l.Logger.With("source", "fx"),
		}
	}
	return &fxevent.NopLogger
}
