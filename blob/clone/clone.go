package clone

import (
	"context"
	"fmt"
	"io"

	gc_blob "gocloud.dev/blob"
)

// Clone copies 'source_uri' from 'source_bucket' to 'target_uri' in 'target_bucket'.
func Clone(ctx context.Context, source_bucket *gc_blob.Bucket, source_uri string, target_bucket *gc_blob.Bucket, target_uri string) error {

	source_r, err := source_bucket.NewReader(ctx, source_uri, nil)

	if err != nil {
		return fmt.Errorf("Failed to create new reader for %s, %w", source_uri, err)
	}

	defer source_r.Close()

	target_wr, err := target_bucket.NewWriter(ctx, target_uri, nil)

	if err != nil {
		return fmt.Errorf("Failed to create new writer for %s, %w", target_uri, err)
	}

	_, err = io.Copy(target_wr, source_r)

	if err != nil {
		target_bucket.Delete(ctx, target_uri)
		return fmt.Errorf("Failed to copy %s to %s, %w", source_uri, target_uri, err)
	}

	err = target_wr.Close()

	if err != nil {
		return fmt.Errorf("Failed to close writer for %s, %w", target_uri, err)
	}

	return nil
}
