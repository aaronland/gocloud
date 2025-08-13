package read

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var bucket_uri string
var key string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("read")

	fs.StringVar(&bucket_uri, "bucket-uri", "", "...")
	fs.StringVar(&key, "key", "", "...")

	return fs
}
