GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
PLATFORM=linux
PLATFORM_WINDOWS=windows-8.1
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
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME_PREFIX)$(PLATFORM_WINDOWS)-$(ARCH).exe

