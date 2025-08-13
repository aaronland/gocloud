package copyuri

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/aaronland/gocloud/blob/bucket"
	"github.com/aaronland/gocloud/blob/copy"
	"github.com/aaronland/gocloud/blob/s3"
	"github.com/aws/aws-lambda-go/lambda"
	aws_s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"gocloud.dev/blob"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	bucket, err := bucket.OpenBucket(ctx, opts.TargetBucketURI)

	if err != nil {
		return fmt.Errorf("Failed to open target bucket, %w", err)
	}

	defer bucket.Close()

	copy_opts := &copy.CopyURLOptions{
		Bucket:       bucket,
		Filename:     opts.Filename,
		ShowProgress: opts.ShowProgress,
	}

	if opts.ACL != "" {

		acl, err := s3.StringACLToObjectCannedACL(opts.ACL)

		if err != nil {
			return fmt.Errorf("Failed to derive ACL object, %w", err)
		}

		before := func(asFunc func(interface{}) bool) error {

			req := &aws_s3.PutObjectInput{}
			ok := asFunc(&req)

			if ok {
				req.ACL = acl
			}

			return nil
		}

		copy_opts.WriterOptions = &blob.WriterOptions{
			BeforeWrite: before,
		}
	}

	do_copy := func(ctx context.Context, copy_opts *copy.CopyURLOptions, source string) error {

		if copy_opts.Filename == "" {
			opts.Filename = filepath.Base(source)
		}

		return copy.CopyURLStringToBucket(ctx, copy_opts, source)
	}

	switch mode {
	case "cli":

		err := do_copy(ctx, copy_opts, opts.SourceURI)

		if err != nil {
			return fmt.Errorf("Failed to copy URI to bucket, %w", err)
		}

	case "lambda":

		handler := func(ctx context.Context, source string) error {
			return do_copy(ctx, copy_opts, source)
		}

		lambda.Start(handler)

	default:
		return fmt.Errorf("Invalid mode")
	}

	return nil
}
