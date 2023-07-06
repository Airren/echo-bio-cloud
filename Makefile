serviceName=echo-bio-cloud


.PHONY: run

run: generate
	go run main.go

.PHONY: build_mac
build_mac: generate
	GOOS=darwin GOARCH=arm64 go build -o ./bin main.go

.PHONY: build_linux
build_linux: generate
	GOOS=linux GOARCH=amd64 go build -o ./bin main.go

.PHONY: build_image
build_image: generate
	docker build --network=host -f Dockerfile ./ -t $(serviceName):latest
.PHONY: generate
generate:
	go generate


.PHONY: help
help:
	@echo 'run              - run project locally'
	@echo 'build_mac        - build binary'
	@echo 'build_linux      - build linux binary'
	@echo 'build_image      - build docker images'
	@echo 'generate         - generate swag file'
