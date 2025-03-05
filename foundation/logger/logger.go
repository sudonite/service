// Package logger provides a convenience function to constructing a logger for use.
// This is required not just for apllications but for testing.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New constructs a Sugared Logger that writes to stdout and provides human-readable timestamps.
func New(service string, outputhPaths ...string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": service,
	}

	config.OutputPaths = []string{"stdout"}
	if outputhPaths != nil {
		config.OutputPaths = outputhPaths
	}

	log, err := config.Build(zap.WithCaller(true))
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
