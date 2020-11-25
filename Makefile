.PHONY: test cover

test:
	go test -race

cover:
	go test -race  -coverprofile=coverage.txt -covermode=atomic
