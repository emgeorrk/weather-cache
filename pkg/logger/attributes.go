package logger

import (
	"log/slog"
	"weather-cache/internal/constants"
)

// Err returns a slog.Attr with the error key and the error message as a string value.
func (l Logger) Err(err error) slog.Attr {
	return slog.Attr{
		Key:   constants.LogErrorKey,
		Value: slog.StringValue(err.Error()),
	}
}

func (l Logger) RequestID(reqID string) slog.Attr {
	return slog.Attr{
		Key:   constants.LogRequestIDKey,
		Value: slog.StringValue(reqID),
	}
}
