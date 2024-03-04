package pmutils

import (
	"log/slog"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SetupLogging() {
	slogOpts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slogOpts))
	slog.SetDefault(logger)
}
