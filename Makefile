.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt: fmt
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

inst: inst
	go install && rm ./intro ./goscrapy
