package runtimevar

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	Timeout int
	Vars    []string
}

func OptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "RUNTIMEVAR")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive flags from environment variables, %w", err)
	}

	vars := fs.Args()

	opts := &RunOptions{
		Timeout: timeout,
		Vars:    vars,
	}

	return opts, nil
}
