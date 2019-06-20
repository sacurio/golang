GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

HELLO_PROG=hello_world
HELLO_UNIX=$(HELLO_PROG)_unix

ARRAY_PROG=array_examples

#all: test build

all: build

build: hello array

hello:
		@echo " > Building $(HELLO_PROG) code..."
		$(GOBUILD) -o $(HELLO_PROG) -v $(HELLO_PROG).go

array:
		@echo " > Building $(ARRAY_PROG) code..."
		$(GOBUILD) -o $(ARRAY_PROG) -v $(ARRAY_PROG).go
test:
		@echo " > Testing..."
		$(GOTEST) -v ./...
clean:
		@echo " > Cleaning build cache"
		$(GOCLEAN)
		rm -f $(HELLO_PROG) $(ARRAY_PROG)
		rm -f $(HELLO_UNIX)
run:
		@echo " > Running code..."
		./$(ARRAY_PROG)

array-run: clean array run


#Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(HELLO_UNIX) -v
docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /home/sham/sham/repos/golang golang:latest go build -o "$(BUILD_UNIX)" -v
