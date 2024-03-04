package main

import (
	"log/slog"

	"pmutils"
)

func main() {
	pmutils.SetupLogging()
	slog.Info("Hello, World!")

}
