all: bins

clean:
	go clean -x

bins:
	go build

tests: bins
	go test -v -race -covermode=atomic -coverprofile=morph.coverprofile

coverage: tests
	go tool cover -html=morph.coverprofile

benchmarks: bins
	go test -C internal -run XXX -bench=.
