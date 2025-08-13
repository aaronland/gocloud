package runtimevar

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var timeout int

func DefaultFlagSet() *flag.FlagSet {
	fs := flagset.NewFlagSet("publish")
	fs.IntVar(&timeout, "timeout", 0, "The maximum number of second in which a variable can be resolved. If 0 no timeout is applied.")
	return fs
}
