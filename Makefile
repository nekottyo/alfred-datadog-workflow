.PHONY: build tidy lint

PKG_NAME = alfred-datadog-workflow
RELEASE_DIR := release

build: get tidy
	@go build -o ${PKG_NAME} .

get:
	@go get -v -t -d ./...

clean: clean-release
	@go clean

tidy:
	@go mod tidy

lint:
	@golangci-lint run

action-release: get
	@mkdir -p ${RELEASE_DIR}
	@go build -o ${RELEASE_DIR}/${PKG_NAME} .
	@cp info.plist icon.png LICENSE README.md ${RELEASE_DIR}

clean-release:
	@rm -rf ${RELEASE_DIR}

