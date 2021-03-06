all: test vet fmt lint build

test:
    go test ./...

vet:
    go vet ./...

fmt:
    go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
    test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
    go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
    go build -o bin/alpha ./cmd/alpha
    go build -o bin/bravo ./cmd/bravo
    go build -o bin/charlie ./cmd/charlie
    go build -o bin/delta ./cmd/delta
    go build -o bin/echo ./cmd/echo
    go build -o bin/foxtrot ./cmd/foxtrot
    go build -o bin/golf ./cmd/golf
    go build -o bin/hotel ./cmd/hotel
    go build -o bin/india ./cmd/india
    go build -o bin/juliet ./cmd/juliet
    go build -o bin/kilo ./cmd/kilo
    go build -o bin/lima ./cmd/lima
    go build -o bin/mike ./cmd/mike
    go build -o bin/november ./cmd/november
    go build -o bin/oscar ./cmd/oscar