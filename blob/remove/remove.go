package remove

import (
	"context"
	"io"
	"log/slog"
	"strings"

	"gocloud.dev/blob"
)

// RemoveTreeOptions defines configuration settings for remove bucket trees.
type RemoveTreeOptions struct {
	// Iterate through each object in a bucket tree but do not remove anything.
	DryRun bool
	// Continue removal even if individual object deletion fails.
	Forgiving bool
}

// RemoveTree will remove 'uri' and all its contents from bucket 'b' with default options.
func RemoveTree(ctx context.Context, b *blob.Bucket, uri string) error {
	opts := &RemoveTreeOptions{}
	return RemoveTreeWithOptions(ctx, b, uri, opts)
}

// RemoveTreeWithOptions will remove 'uri' and all its contents from bucket 'b' configure using 'opts'.
func RemoveTreeWithOptions(ctx context.Context, b *blob.Bucket, uri string, opts *RemoveTreeOptions) error {

	var removeTree func(context.Context, *blob.Bucket, string) error

	removeTree = func(ctx context.Context, b *blob.Bucket, prefix string) error {

		logger := slog.Default()
		logger = logger.With("uri", uri)
		logger = logger.With("prefix", prefix)

		iter := b.List(&blob.ListOptions{
			Delimiter: "/",
			Prefix:    prefix,
		})

		for {
			obj, err := iter.Next(ctx)

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			if obj.IsDir {

				err = removeTree(ctx, b, obj.Key)

				if err != nil {
					logger.Error("Failed to remove tree", "path", obj.Key)
					return err
				}

			}

			// trailing slashes confuse Go Cloud...

			path := strings.TrimRight(obj.Key, "/")

			if opts.DryRun {
				logger.Debug("Delete object here (dryrun)", "path", path)
			} else {
				err = b.Delete(ctx, path)

				if err != nil {
					logger.Error("Failed to remove object", "path", path, "error", err)
				}

				if !opts.Forgiving {
					return err
				}
			}
		}

		return nil
	}

	return removeTree(ctx, b, uri)
}
