package runtimevar

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/aaronland/gocloud/runtimevar"
)

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	opts, err := OptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *slog.Logger) error {

	if opts.Timeout > 0 {

		c, cancel := context.WithTimeout(ctx, time.Duration(opts.Timeout)*time.Second)
		defer cancel()

		ctx = c
	}

	for _, uri := range opts.Vars {

		str_var, err := runtimevar.StringVar(ctx, uri)

		if err != nil {
			return fmt.Errorf("Failed to derive variable, %w", err)
		}

		fmt.Println(str_var)
	}

	return nil
}
