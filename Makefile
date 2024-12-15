
all: fmt vet test

build: test


fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -v --cover ./...


