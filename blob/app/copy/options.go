package copy

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
)

type RunOptions struct {
	Source        string
	Target        string
	ACL           string
	WriterOptions *blob.WriterOptions
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	opts := &RunOptions{
		Source:        source,
		Target:        target,
		WriterOptions: new(blob.WriterOptions),
		ACL:           str_acl,
	}

	return opts, nil
}
