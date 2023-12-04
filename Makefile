.PHONY: help
help: ### Shows this help
	@grep -E '^[0-9a-zA-Z_-]+:' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?### "}; {printf "\033[32;1m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ### Runs this stuff
	go test -coverprofile coverage.out ./...
	go tool cover -func coverage.out
	go tool cover -html coverage.out -o coverage.html
	open coverage.html

.PHONY: example
example: ### Runs this stuff
	cd ./example && go run .
