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

func StrToBool(s string) bool {
	return s != ""
}

func SetupLogging() {
	logCode := slog.LevelInfo
	if StrToBool(GetEnv("DEBUG", "")) {
		logCode = slog.LevelDebug
	}

	slogOpts := slog.HandlerOptions{
		Level: logCode,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slogOpts))
	slog.SetDefault(logger)
}

func ContainsEmptyString(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}
