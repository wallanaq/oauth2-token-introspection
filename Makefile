APP_NAME     := oauth2-token-introspection
APP_VERSION  := 1.0.0
BIN_DIR      := bin/
DOCKER_IMAGE := wallanaq/$(APP_NAME):$(APP_VERSION)

.PHONY: all build clean run build-image docker-run

## Compile the application
build:
	@echo ">> Building..."
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/api

## Run the application locally
run: build
	@echo ">> Running locally..."
	./$(BIN_DIR)/$(APP_NAME)

## Clean up binaries and temporary files
clean:
	@echo ">> Cleaning up..."
	rm -rf $(BIN_DIR)

## Build the Docker image
build-image:
	@echo ">> Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

## Run Docker container (detached)
docker-run:
	docker run -d -p 8080:8080 --name $(APP_NAME) $(DOCKER_IMAGE)

## Stop and remove Docker container
docker-clean:
	docker rm -f $(APP_NAME) || true
	docker rmi -f $(DOCKER_IMAGE) || true

## Rebuild image and run container
rebuild: docker-clean build-image docker-run
