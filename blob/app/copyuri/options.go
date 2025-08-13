package copyuri

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	TargetBucketURI string
	SourceURI       string
	Filename        string
	ShowProgress    bool
	Mode            string
	ACL             string
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "BLOB")

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment, %w", err)
	}

	opts := &RunOptions{
		TargetBucketURI: target_uri,
		SourceURI:       source_uri,
		Filename:        filename,
		ShowProgress:    show_progress,
		Mode:            mode,
		ACL:             str_acl,
	}

	return opts, nil
}
