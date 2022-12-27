BINDIR=bin
BINNAME=cource-registration-system-backend
BINPATH=$(BINDIR)/$(BINNAME)

.PHONY: build
build:
	go build -o $(BINPATH) .

.PHONY: clean
clean:
	rm $(BINPATH) && go clean