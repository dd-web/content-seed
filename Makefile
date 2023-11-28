.PHONY: test

download:
	go mod download

test: download ## Run all tests
	go test -v

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

bench: download ## Run all benchmarks
	go test -bench=. -benchmem