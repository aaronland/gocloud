package walk

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
)

func TestWalkBucket(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive current path, %v", err)
	}

	bucket_uri := fmt.Sprintf("file://%s", abs_path)

	bucket, err := blob.OpenBucket(ctx, bucket_uri)

	if err != nil {
		t.Fatalf("Failed to open bucket for %s, %v", bucket_uri, err)
	}

	defer bucket.Close()

	files := make([]string, 0)

	cb := func(ctx context.Context, obj *blob.ListObject) error {
		files = append(files, obj.Key)
		return nil
	}

	err = WalkBucket(ctx, bucket, cb)

	if err != nil {
		t.Fatalf("Failed to walk bucket, %v", err)
	}

	str_files := strings.Join(files, " ")
	expected := "walk.go walk_test.go"

	if str_files != expected {
		t.Fatalf("Unexpected files: %s", str_files)
	}
}
