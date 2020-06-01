.PHONY: build tidy lint

PKG_NAME = alfred-datadog-workflow
RELEASE_DIR := release
BUILD_CMD := CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'

build: get tidy
	@${BUILD_CMD} -o ${PKG_NAME} ${LDFLAGS} .

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
	@${BUILD_CMD} ${RELEASE_DIR}/${PKG_NAME}  ${LDFLAGS}  .
	@cp info.plist icon.png LICENSE README.md ${RELEASE_DIR}

clean-release:
	@rm -rf ${RELEASE_DIR}

