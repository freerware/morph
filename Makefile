all: bins

clean:
	go clean -x

bins:
	go build

test: bins
	go test -v -race -covermode=atomic -coverprofile=coverage.out
