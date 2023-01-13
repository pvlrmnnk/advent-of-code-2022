.PHONY: lint
lint:
	@golangci-lint run ./solution/...

.PHONY: test
test:
	@go test ./solution/...

.PHONY: solution
solution:
	@cookiecutter -o ./solution .cookiecutter/solution
