all: mod test vet fmt build run

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

mod:
	go mod tidy
	go mod vendor

build:
	go build -o bin/gitcloak main.go

run:
	./bin/gitcloak

install:
	go install -v ./...

clean:
	rm -rf .gitcloak
	rm -f bin/*

