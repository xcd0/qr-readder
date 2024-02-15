VERSION       := 1.0.0
REVISION      := `git rev-parse --short HEAD`
FLAGS_VERSION := -X main.version=$(VERSION) -X main.revision=$(shell git rev-parse --short HEAD)
FLAG := -ldflags='-X main.version=$(VERSION) -X main.revision='$(REVISION)' -s -w -extldflags="-static" -buildid=' -a -tags netgo -installsuffix -trimpath 
#FLAG := -ldflags='-X main.version=$(VERSION) -X main.revision='$(REVISION)' -s -w -extldflags="-static" -H windowsgui -buildid=' -a -tags netgo -installsuffix -trimpath 


all:
	cat makefile

fmt:
	goimports -w *.go
	gofmt -w *.go

gen:
	go generate

build:
	go build $(FLAG)

upx:
	upx --lzma `cat go.mod | grep module | awk '{print $$2}'`
upx-exe:
	upx --lzma `cat go.mod | grep module | awk '{print $$2}'`.exe

release:
	make build
	make upx
	GOOS=windows make build
	make upx-exe

get:
	go mod tidy && go get

