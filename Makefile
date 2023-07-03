.PHONY: test
test: ### Runs this stuff
	go test -coverprofile cover.out ./...
	go tool cover -func cover.out
	go tool cover -html cover.out -o cover.html
	open cover.html

.PHONY: example
example: ### Runs this stuff
	cd ./example && go run .
