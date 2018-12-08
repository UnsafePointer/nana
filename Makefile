GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
PLATFORM=linux
PLATFORM_WINDOWS=windows-8.1
PLATFORM_DARWIN=darwin-10.11
ARCH=amd64
BINARY_NAME_PREFIX=nana-
BINARY_NAME=$(BINARY_NAME_PREFIX)$(PLATFORM)-$(ARCH)

all: build test
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME_PREFIX)*

# Cross-compiling

build-windows:
	xgo --targets=$(PLATFORM_WINDOWS)/$(ARCH) github.com/Ruenzuo/nana
build-darwin:
	xgo --targets=$(PLATFORM_DARWIN)/$(ARCH) github.com/Ruenzuo/nana

