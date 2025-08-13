package clone

import (
	"context"
	"fmt"
	"io"
	"io/fs"

	"gocloud.dev/blob"
)

// CloneFS will opy the contents of `to_clone' to 'target_bucket'.
func CloneFS(ctx context.Context, to_clone fs.FS, target_bucket *blob.Bucket, wr_opts *blob.WriterOptions) error {

	walk_func := func(path string, d fs.DirEntry, err error) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		r, err := to_clone.Open(path)

		if err != nil {
			return fmt.Errorf("Failed to open %s for reading, %w", path, err)
		}

		defer r.Close()

		wr, err := target_bucket.NewWriter(ctx, path, wr_opts)

		if err != nil {
			return fmt.Errorf("Failed to open %s for writing, %w", path, err)
		}

		_, err = io.Copy(wr, r)

		if err != nil {
			return fmt.Errorf("Failed to write %s to target bucket, %w", path, err)
		}

		err = wr.Close()

		if err != nil {
			return fmt.Errorf("Failed to close %s after writing, %w", path, err)
		}

		return nil
	}

	return fs.WalkDir(to_clone, ".", walk_func)
}
