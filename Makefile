.PHONY: build tidy lint

PKG_NAME = alfred-datadog-workflow

build: tidy
	@go build -o ${PKG_NAME} .

clean:
	@go clean

tidy:
	@go mod tidy

lint:
	@golangci-lint run
