package bucket

import (
	"testing"
)

func TestParseURI(t *testing.T) {

	tests := map[string][2]string{
		"/usr/local/example.jpg": [2]string{
			"file:///?prefix=usr%2Flocal%2F",
			"example.jpg",
		},
		"file:///usr/local/example.jpg": [2]string{
			"file:///usr/local/", "example.jpg",
		},
		"s3blob://example/folder/example.txt?region=us-east-1&credentials=session": [2]string{
			"s3blob://example?credentials=session&prefix=folder%2F&region=us-east-1", "example.txt",
		},
		"s3://example/folder/example.txt?region=us-east-1&credentials=session": [2]string{
			"s3://example?credentials=session&prefix=folder%2F&region=us-east-1", "example.txt",
		},
		"s3://example/example.txt?prefix=folder/&region=us-east-1&credentials=session": [2]string{
			"s3://example?credentials=session&prefix=folder%2F&region=us-east-1", "example.txt",
		},
		"s3://example/test/example.txt?prefix=folder/&region=us-east-1&credentials=session": [2]string{
			"s3://example?credentials=session&prefix=folder%2Ftest%2F&region=us-east-1", "example.txt",
		},
	}

	for uri, expected := range tests {

		bucket_uri, bucket_key, err := ParseURI(uri)

		if err != nil {
			t.Fatalf("Failed to parse URI '%s', %v", uri, err)
		}

		if bucket_uri != expected[0] {
			t.Fatalf("Unexpected bucket URI for '%s'. Expected '%s' but got '%s'.", uri, expected[0], bucket_uri)
		}

		if bucket_key != expected[1] {
			t.Fatalf("Unexpected bucket key for '%s'. Expected '%s' but got '%s'.", uri, expected[1], bucket_key)
		}
	}
}
