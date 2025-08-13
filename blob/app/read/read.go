package read

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/aaronland/gocloud/blob/bucket"
	"github.com/sfomuseum/go-flags/flagset"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	source_bucket, err := bucket.OpenBucket(ctx, bucket_uri)

	if err != nil {
		return fmt.Errorf("Failed to open source bucket '%s', %v", bucket_uri, err)
	}

	defer source_bucket.Close()

	r, err := source_bucket.NewReader(ctx, key, nil)

	if err != nil {
		return fmt.Errorf("Failed to open %s for reading, %v", key, err)
	}

	defer r.Close()

	_, err = io.Copy(os.Stdout, r)

	if err != nil {
		return fmt.Errorf("Failed to copy %s to STDOUT, %v", key, err)
	}

	return nil
}
