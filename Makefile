.PHONY: test test-report

test:
	go test ./...

test-report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
