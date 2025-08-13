package s3

import (
	"context"
	"flag"
	"io"
	"testing"

	"gocloud.dev/blob"
)

var uri = flag.String("uri", "", "A valid s3blob:// URI.")

func TestOpenBucket(t *testing.T) {

	ctx := context.Background()

	if *uri == "" {
		t.Log("Empty -uri flag. Skipping.")
		t.Skip()
	}

	bucket, err := blob.OpenBucket(ctx, *uri)

	if err != nil {
		t.Fatal(err)
	}

	defer bucket.Close()

	iter := bucket.List(nil)

	for {
		obj, err := iter.Next(ctx)

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		t.Log(obj.Key)
	}
}
