# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
# GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOSEC=$(GOCMD)sec

all: test build inspect
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
inspect:
	gosec ./...
	golint ./...
	go vet ./...
sec:
	$(GOSEC) ./...
test: inspect
	$(GOTEST) ./...
clean: # remove o binanrio gerado 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
# run: # descomentar essa linha para não executar o programa
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

# Basic command
BINARY_NAME=daur