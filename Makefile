PACKAGES = $(shell go list ./...)

all: test

dep:
	@dep ensure

build:
	@go build -o bin/server github.com/s-ichikawa/piql/server

test:
	@go test -v $(PACKAGES)

vet:
	@go vet $(PACKAGES)

coverage:
	@go get -u github.com/go-playground/overalls
	overalls \
		-project github.com/s-ichikawa/piql \
		-covermode atomic \
		-concurrency 8
	@mv overalls.coverprofile coverage.txt


.PHONY: all test vet cover