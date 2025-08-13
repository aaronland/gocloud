package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aaronland/gocloud/runtimevar/app/runtimevar"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := runtimevar.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to run application", "error", err)
		os.Exit(1)
	}
}
