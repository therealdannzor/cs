DIRS_TEST = ./list

test: 
	@go test $(DIRS_TEST) -failfast

lint:
	golangci-lint run -E gofmt -E golint -E gosec --deadline 1m0s $(DIRS_TEST)
