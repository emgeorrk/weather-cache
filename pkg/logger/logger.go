package logger

import (
	"github.com/charmbracelet/log"
	"go.uber.org/fx"
	"log/slog"
	"os"
	"strings"
	"weather-cache/config"
)

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.WithLogger(NewFxLogger),
)

type Logger struct {
	*slog.Logger
}

func (l Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func NewLogger(config *config.Config) Logger {
	level := log.InfoLevel
	switch strings.ToLower(config.Log.Level) {
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warn":
		level = log.WarnLevel
	case "error":
		level = log.ErrorLevel
	}

	formatter := log.TextFormatter
	switch strings.ToLower(config.Log.Formatter) {
	case "json":
		formatter = log.JSONFormatter
	case "text":
		formatter = log.TextFormatter
	case "fmt":
		formatter = log.LogfmtFormatter
	}

	sourceFormat := log.ShortCallerFormatter
	if strings.ToLower(config.Log.SourceFormat) == "long" {
		sourceFormat = log.LongCallerFormatter
	}

	log.Helper()

	opts := log.Options{
		ReportTimestamp: strings.ToLower(config.Log.AddTimestamp) == "yes",
		TimeFormat:      config.Log.TimeFormat,
		Level:           level,
		Prefix:          config.Log.Prefix,
		ReportCaller:    strings.ToLower(config.Log.AddSource) == "yes",
		CallerFormatter: sourceFormat,
		Formatter:       formatter,
	}

	handler := log.NewWithOptions(os.Stderr, opts)

	if config.Log.Label != "" {
		handler = handler.With("label", config.Log.Label)
	}

	logger := slog.New(handler)

	logger.Info("Logger initialized", "log_level", config.Log.Level)

	return Logger{
		Logger: logger,
	}
}
