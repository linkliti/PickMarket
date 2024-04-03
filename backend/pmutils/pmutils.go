package pmutils

import (
	"fmt"
	"io"
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

func SetupLogging(module string) {
	debug := StrToBool(GetEnv("DEBUG", ""))
	logCode := slog.LevelInfo
	if debug {
		logCode = slog.LevelDebug
	}
	slogOpts := slog.HandlerOptions{
		Level: logCode,
	}
	logFile, err := os.OpenFile(module+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file %s for output: %s", module+".log", err)
		panic(err)
	}
	handler := slog.NewJSONHandler(logFile, &slogOpts)
	if debug {
		handler = slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &slogOpts)
	}
	logger := slog.New(handler)
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
