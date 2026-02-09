test:
	go test ./...

coverage:
	go test ./... -coverpkg=./cmd/... -coverprofile=coverage.out
	go tool cover -html=coverage.out

coverage-cli:
	go tool cover -func=coverage.out
