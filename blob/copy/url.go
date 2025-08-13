package copy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/mitchellh/ioprogress"
	"gocloud.dev/blob"
)

type CopyURLOptions struct {
	// The `blob.Bucket` instance where the body of URL requests will be copied to.
	Bucket *blob.Bucket
	// Any option `blob.WriterOptions` to specify when creating a new `blob.Writer` instance.
	WriterOptions *blob.WriterOptions
	// An optional filename for the new file. If empty then basename of the URL's `Path` component will be used.
	Filename string
	// Display a progress meter as a file is being copied in to a bucket
	ShowProgress bool
}

// CopyURLToBucket will parse 'uri' to ensure it is a valid `url.URL` instance and then
// fetch and copy its contents to a file in a `blob.Bucket` instance.
func CopyURLStringToBucket(ctx context.Context, opts *CopyURLOptions, uri string) error {

	u, err := url.Parse(uri)

	if err != nil {
		return fmt.Errorf("Failed to parse %s, %w", uri, err)
	}

	return CopyURLToBucket(ctx, opts, u)
}

// CopyURLToBucket will fetch and copy the contents of 'u' to a file in a `blob.Bucket` instance.
func CopyURLToBucket(ctx context.Context, opts *CopyURLOptions, u *url.URL) error {

	fname := opts.Filename

	if fname == "" {
		fname = filepath.Base(u.Path)
	}

	uri := u.String()

	cl := &http.Client{}

	if u.Scheme == "file" {

		uri_dir := filepath.Dir(u.Path)
		uri_fname := filepath.Base(u.Path)

		tr := &http.Transport{}
		tr.RegisterProtocol("file", http.NewFileTransport(http.Dir(uri_dir)))

		cl = &http.Client{Transport: tr}

		u.Path = fmt.Sprintf("/%s", uri_fname)
		uri = u.String()
	}

	rsp, err := cl.Get(uri)

	if err != nil {
		return fmt.Errorf("Failed to GET %s, %w", uri, err)
	}

	defer rsp.Body.Close()

	wr, err := opts.Bucket.NewWriter(ctx, fname, opts.WriterOptions)

	if err != nil {
		return fmt.Errorf("Failed to create new writer for %s, %w", fname, err)
	}

	var r io.Reader

	if opts.ShowProgress {

		content_length := rsp.Header.Get("Content-Length")
		size, err := strconv.ParseInt(content_length, 10, 64)

		if err != nil {
			return fmt.Errorf("Failed to determine content length for '%s', %w", content_length, err)
		}

		r = &ioprogress.Reader{
			Reader: rsp.Body,
			Size:   size,
		}
	} else {
		r = rsp.Body
	}

	_, err = io.Copy(wr, r)

	if err != nil {
		return fmt.Errorf("Failed to copy data, %w", err)
	}

	err = wr.Close()

	if err != nil {
		return fmt.Errorf("Failed to close %s, %w", fname, err)
	}

	return nil
}
