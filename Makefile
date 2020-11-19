test:
	@go test -v -coverprofile .coverage -race -timeout 30s ./...
	@go tool cover -func .coverage

coverage: test
	@go tool cover -html=.coverage
