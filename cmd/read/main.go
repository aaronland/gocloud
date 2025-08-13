package main

import (
	"context"
	"log"

	_ "github.com/aaronland/gocloud/blob/s3"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
	
	"github.com/aaronland/gocloud/blob/app/read"
)

func main() {

	ctx := context.Background()
	err := read.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
