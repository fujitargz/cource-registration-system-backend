BINDIR=bin
BINNAME=trial-api
BINPATH=$(BINDIR)/$(BINNAME)

.PHONY: build
build:
	CGO_ENABLED=1 CGO_LDGLAGS=' -linkmode external -extldflags "-static"' GOOS=linux go build -o $(BINPATH) ./cmd/$(BINNAME)/main.go

.PHONY: clean
clean:
	rm $(BINPATH) && go clean

.PHONY: run
run:
	go run ./cmd/$(BINNAME)/main.go