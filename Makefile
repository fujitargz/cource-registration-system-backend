BINDIR=bin
BINNAME=trial-api
BINPATH=$(BINDIR)/$(BINNAME)

.PHONY: build
build:
	CGO_ENABLED=1 go build -ldflags '-linkmode external -extldflags "-static"' -o $(BINPATH) ./cmd/$(BINNAME)/main.go

.PHONY: clean
clean:
	rm $(BINPATH) && go clean

.PHONY: run
run:
	docker build -t trial-api:0.1 -f ./build/Dockerfile . && docker run -p 8000:8000 trial-api:0.1