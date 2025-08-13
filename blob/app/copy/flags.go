package copy

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var source string
var target string

var str_acl string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("copy")

	fs.StringVar(&source, "source", "", "The fully-qualified URI for the source file (this will be parsed in to a gocloud.dev/blob.Bucket URI and filename).")
	fs.StringVar(&target, "target", "", "The fully-qualified URI for the target file (this will be parsed in to a gocloud.dev/blob.Bucket URI and filename).")
	fs.StringVar(&str_acl, "acl", "", "An optional AWS S3 ACL to assign to the file being copied.")

	return fs
}
