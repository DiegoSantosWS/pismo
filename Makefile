# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
# GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -test.short
GOTEST_FUNC=$(GOCMD) test
GOSEC=$(GOCMD)sec

all: test build inspect
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
inspect:
	gosec ./...
	golint ./... -v
	go vet ./...
sec:
	$(GOSEC) ./...
test: inspect
	$(GOTEST) ./... -v
test-integration:
	$(GOTEST_FUNC) ./... -v
build:
	docker build . -t pismo:latest
clean: # remove o binanrio gerado 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
# run: # descomentar essa linha para não executar o programa
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

# Basic command
BINARY_NAME=pismo