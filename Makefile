build:
	CGO_ENABLED=0 go build -o cal_service

GO_TEST_PATHS := $(shell command go list ./... | grep -v "vendor")
test:
	go test  $(GO_TEST_PATHS)

clean:
	rm -f cal_service

.DEFAULT_GOAL := build