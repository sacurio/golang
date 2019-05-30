GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
		@echo " > Building code..."
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		@echo " > Testing..."
		$(GOTEST) -v ./...
clean:
		@echo " > Cleaning build cache"
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		@echo " > Running code..."
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

#Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /home/sham/sham/repos/golang golang:latest go build -o "$(BUILD_UNIX)" -v