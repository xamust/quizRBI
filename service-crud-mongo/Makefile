init:
		go mod tidy
.PHONY: init
build:
		make init
		go build -o build/service-crud-mongo -v ./cmd/app
.PHONY: build
run:
		make build
		./build/service-crud-mongo
.PHONY: run
pretest: ### get linters
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go get -u golang.org/x/lint/golint@latest
.PHONY: pretest
test:
	make pretest
	golint ./...
	go vet ./...
	staticcheck ./...
.PHONY: test

.DEFAULT_GOAL := run