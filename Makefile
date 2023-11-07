.PHONY: build-macos

GIT_COMMIT := $(shell git rev-list -1 HEAD)
DATE := $(shell date)
VERSION := $(shell git describe --tags --abbrev=0)

build-macos:
	GOOS=darwin
	GOARCH=arm64
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.buildDate=$(DATE)' -X 'main.buildVersion=$(VERSION)'" -o koche

build-windows:
	GOOS=windows
	GOARCH=amd64
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.buildDate=$(DATE)' -X 'main.buildVersion=$(VERSION)'" -o koche.exe .
