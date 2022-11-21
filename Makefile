serviceName=echo-bio-cloud


.PHONY: run
run:
	go run main.go

.PHONY: build_mac
build_mac:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/exporter-demo main.go

.PHONY: build_linux
build_linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/exporter-demo main.go

.PHONY: build_image
build_image:
	docker build --network=host -f Dockerfile ./ -t $(serviceName):latest


.PHONY: help
help:
	@echo 'run              - run project locally'
	@echo 'build_mac        - build binary'
	@echo 'build_linux      - build linux binary'
	@echo 'build_image      - build docker images'