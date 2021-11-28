GOCMD=go
GOFMT=$(GOCMD) fmt
GOVETT=$(GOCMD) vet
GOBUILD=$(GOCMD) build

all:
	make validate
	make build

validate:
	$(GOFMT) ./...

build:
	$(GOBUILD) main.go

