package logger

import (
	"context"
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

// RequestID returns a slog.Attr with the request ID key and the request ID as a string value.
func (l Logger) RequestID(ctx context.Context) slog.Attr {
	return slog.Attr{
		Key:   constants.LogRequestIDKey,
		Value: slog.AnyValue(ctx.Value(constants.RequestID)),
	}
}

// String returns a slog.Attr with the key and value as a string value.
func (l Logger) String(key, value string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.StringValue(value),
	}
}

func (l Logger) With(attr slog.Attr) Logger {
	return Logger{
		Logger: l.Logger.With(attr),
	}
}
