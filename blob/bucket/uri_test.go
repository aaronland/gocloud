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
