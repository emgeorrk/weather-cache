package logger

import (
	"github.com/charmbracelet/log"
	"go.uber.org/fx"
	"os"
	"strings"
	"weather-cache/config"
)

var Module = fx.Options(
	fx.Provide(New),
)

type Logger struct {
	*log.Logger
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func New(config config.Config) Logger {
	level := log.InfoLevel
	switch strings.ToLower(config.LogLevel) {
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
	switch strings.ToLower(config.LogFormatter) {
	case "json":
		formatter = log.JSONFormatter
	case "text":
		formatter = log.TextFormatter
	case "fmt":
		formatter = log.LogfmtFormatter
	}

	sourceFormat := log.ShortCallerFormatter
	if strings.ToLower(config.LogSourceFormat) == "long" {
		sourceFormat = log.LongCallerFormatter
	}

	log.Helper()

	opts := log.Options{
		ReportTimestamp: strings.ToLower(config.LogAddTimeStamp) == "yes",
		TimeFormat:      config.LogTimeFormat,
		Level:           level,
		Prefix:          config.LogPrefix,
		ReportCaller:    strings.ToLower(config.LogAddSource) == "yes",
		CallerFormatter: sourceFormat,
		Formatter:       formatter,
	}

	logger := log.NewWithOptions(os.Stderr, opts)

	if config.LogLabel != "" {
		logger = logger.With("label", config.LogLabel)
	}

	logger.Info("Logger initialized", "level", config.LogLevel)

	return Logger{
		Logger: logger,
	}
}
