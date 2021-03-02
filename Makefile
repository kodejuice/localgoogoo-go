GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=localgoogoo-go
BINARY_UNIX=$(BINARY_NAME)_unix
# 
GOTEST=richgo test
GOLINT=golangci-lint
#
XC_OS="linux darwin windows"
XC_ARCH="amd64"
XC_PARALLEL="2"
#
SRC=$(shell find . -name "*.go")
BIN="./bin"


ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif


.PHONY: fmt lint test build install_deps clean

default: fmt test build

fmt:
	$(info ###################### checking formatting ######################)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ###################### running tests ############################)
	$(GOLINT) run -v

test: install_deps lint
	$(GOTEST) -v github.com/kodejuice/localgoogoo-go/search
	$(GOTEST) -v github.com/kodejuice/localgoogoo-go/crawler

build:
	mkdir -p $(BIN)
	# cross compilation
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;

install_deps:
	$(info ###################### downloading dependencies #################)
	go get -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BIN)
