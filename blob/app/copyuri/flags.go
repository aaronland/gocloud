package copyuri

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var source_uri string
var filename string
var target_uri string
var show_progress bool
var mode string

var str_acl string

// var acl string
// var part_size int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("copyuri")

	fs.StringVar(&source_uri, "source-uri", "", "The URI of the file to copy.")
	fs.StringVar(&target_uri, "target-uri", "", "A valid gocloud.dev/blob.Bucket URI.")

	fs.StringVar(&filename, "filename", "", "The final filename of the file to copy. If empty the basename of the `-source-uri` flag value will be used.")

	fs.StringVar(&str_acl, "acl", "", "...")

	// fs.StringVar(&acl, "acl", "", "An optional AWS S3 ACL to assign to the file being copied.")
	// fs.Int64Var(&part_size, "part-size", 0, "The buffer size (in bytes) to use when buffering data into chunks and sending them as parts to S3. If 0 the default value for the `aws/aws-sdk-go/service/s3/s3manager` package will be used.")

	fs.BoolVar(&show_progress, "show-progress", false, "Show copy progress.")
	fs.StringVar(&mode, "mode", "cli", "Valid options are: cli, lambda.")

	return fs
}
