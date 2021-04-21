ifeq ($(GOPATH),)
GOPATH := $(HOME)/go
endif

all: test lint vet build

clean:
	@rm -f bin/*
	@go clean -testcache

build: serialnum

serialnum:
	@go mod tidy
	@cd cmd/serialnum && go build -o ../../bin/serialnum

check: lint vet staticcheck revive

lint:
	@revive -exclude ./third_party/... ./... 

vet:
	@go vet ./...

staticcheck:
	@staticcheck ./...

test:
	@go test ./...

test_verbose:
	@go test ./... -v

test_race:
	@go test ./... -race

test_all: test_cover test_race
