# Parameters
MAIN_FILE=main.go
MAIN_PATH=cmd/search_server
SERVER=search_server
DOCKER_FILE=Dockerfile
DOCKER_TAG=search_server

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o $(MAIN_PATH)/$(SERVER) $(MAIN_PATH)/$(MAIN_FILE)

build-linux:
	GOARCH=amd64 GOOS=linux go build -o $(MAIN_PATH)/$(SERVER) $(MAIN_PATH)/$(MAIN_FILE)

build-docker:
	docker build -t $(DOCKER_TAG) -f $(MAIN_PATH)/$(DOCKER_FILE) .

test:
	go test -race -v ./...