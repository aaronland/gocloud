# gocloud-blob

Opinionated methods and tools for working with gocloud.dev/blob instances.

## Documentation

Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/copy-uri cmd/copy-uri/main.go
go build -mod vendor -ldflags="-s -w" -o bin/copy cmd/copy/main.go
go build -mod vendor -ldflags="-s -w" -o bin/read cmd/read/main.go
```

### copy

```
$> ./bin/copy -h
  -acl string
    	An optional AWS S3 ACL to assign to the file being copied.
  -source string
    	The fully-qualified URI for the source file (this will be parsed in to a gocloud.dev/blob.Bucket URI and filename).
  -target string
    	The fully-qualified URI for the target file (this will be parsed in to a gocloud.dev/blob.Bucket URI and filename).
```

#### Example

For example:

```
$> ./bin/copy -source file:///usr/local/src/gocloud-blob/README.md -target mem:///test.txt
```

### copy-uri

```
$> ./bin/copy-uri -h
  -acl string
    	An optional AWS S3 ACL to assign to the file being copied.
  -filename -source-uri
    	The final filename of the file to copy. If empty the basename of the -source-uri flag value will be used.
  -mode string
    	Valid options are: cli, lambda. (default "cli")
  -show-progress
    	Show copy progress.
  -source-uri string
    	The URI of the file to copy.
  -target-uri string
    	A valid gocloud.dev/blob.Bucket URI.
```

#### Example

For example:

```
$> ./bin/copy-uri \
	-source-uri https://static.sfomuseum.org/media/189/736/720/3/1897367203_lt3HuJ5ALbY4SDoxTd4oi7abOF7gQZKM_c.jpg \
	-target-uri cwd://

$> ll 1897367203_lt3HuJ5ALbY4SDoxTd4oi7abOF7gQZKM_c.jpg 
-rw-r--r--  1 user  staff  342776 Apr 12 17:00 1897367203_lt3HuJ5ALbY4SDoxTd4oi7abOF7gQZKM_c.jpg
```

#### Lambda

The `copy-uri` can also be run as a Lambda function.

```
$> make lambda
if test -f bootstrap; then rm -f bootstrap; fi
if test -f copy_uri.zip; then rm -f copy_uri.zip; fi
GOARCH=arm64 GOOS=linux go build -mod vendor -ldflags="" -tags lambda.norpc -o bootstrap cmd/copy-uri/main.go
zip copy_uri.zip bootstrap
  adding: bootstrap (deflated 62%)
rm -f bootstrap
```

Install and configure the Lambda function as necessary. The following environment variables need to be configured:

| Key | Value | Notes |
| --- | --- | --- |
| BLOB_MODE | lambda | |
| BLOB_TARGET_URI | | A valid `gocloud.dev/blob` URI. For example `s3blob://{S3_BUCKET}?region={S3_REGION}&prefix={S3_PREFIX}/&credentials={CREDENTIALS}`

If you are using the `s3blob://` URI scheme `{CREDENTIALS}` is expected to be a [aaronland/go-aws-session](https://github.com/aaronland/go-aws-session) credentials string:


| Label | Description |
| --- | --- |
| `anon:` | Empty or anonymous credentials. |
| `env:` | Read credentials from AWS defined environment variables. |
| `iam:` | Assume AWS IAM credentials are in effect. |
| `sts:{ARN}` | Assume the role defined by `{ARN}` using STS credentials. |
| `{AWS_PROFILE_NAME}` | This this profile from the default AWS credentials location. |
| `{AWS_CREDENTIALS_PATH}:{AWS_PROFILE_NAME}` | This this profile from a user-defined AWS credentials location. |

Environment variables are mapped to command line flags as followed. For any given command line flag it is:

* Upper-cased
* Spaces are replaced with "_" characters
* The final string is prepended with "BLOB_"

For example the `-target-uri` flag becomes the `BLOB_TARGET_URI` environment variable.

#### Docker

```
$> make docker
docker buildx build --platform=linux/amd64 -t gocloud-blob .
[+] Building 36.1s (14/14) FINISHED                                                                                                                              docker:deskt         
View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/jvr88erchuxoddhbwdhwv7xap
```

```
$> docker run --platform=linux/amd64 gocloud-blob \
	/usr/local/bin/copy-uri \
	-source-uri https://static.sfomuseum.org/media/189/736/720/3/1897367203_lt3HuJ5ALbY4SDoxTd4oi7abOF7gQZKM_c.jpg \
	-target-uri file:///tmp/
```

## See also

* https://pkg.go.dev/gocloud.dev/blob