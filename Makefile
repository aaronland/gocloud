GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/copy cmd/copy/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/copy-uri cmd/copy-uri/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/read cmd/read/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/runtimevar cmd/runtimevar/main.go
