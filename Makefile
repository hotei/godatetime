# Makefile for godatetime program

all:
	go build

install:
	go build
	go tool vet .
	go tool vet -shadow .
	gofmt -w *.go
	cp godatetime ~/bin
	cp README-godatetime.md README.md
