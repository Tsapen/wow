.PHONY: test run-server run-client

export WOW_ROOT_DIR = $(shell pwd)
export WOW_SERVER_CONFIG = /configs/server_config.json
export WOW_CLIENT_CONFIG = /configs/client_config.json
export PORT = $(shell grep -o '"address": "[^"]*"' $(WOW_ROOT_DIR)$(WOW_SERVER_CONFIG) | cut -d ':' -f 3 | cut -d '"' -f 1)

export SERVER_IMAGE = word-of-wisdom-server
export CLIENT_IMAGE = word-of-wisdom-client

install-deps:
	go install github.com/vektra/mockery/v2

generate: install-deps
	go generate ./...

build-test-image:
	docker build -t word-of-wisdom-test -f $(WOW_ROOT_DIR)/deployment/test/Dockerfile .

run-tests:
	docker run --rm word-of-wisdom-test

# Run tests inside a Docker container
test: build-test-image run-tests

# Build and run the server in a Docker container
run-server: stop-server
	docker network rm wow-network 2>/dev/null || true
	docker network create wow-network
	docker build -t $(SERVER_IMAGE) --build-arg WOW_ROOT_DIR=$(WOW_ROOT_DIR) --build-arg WOW_SERVER_CONFIG=$(WOW_SERVER_CONFIG) -f $(WOW_ROOT_DIR)/deployment/server/Dockerfile .
	docker run -d -p $(PORT):$(PORT) --name $(SERVER_IMAGE) --network wow-network $(SERVER_IMAGE)

# Build and run the client in a Docker container
run-client: stop-client
	docker build -t $(CLIENT_IMAGE) --build-arg WOW_ROOT_DIR=$(WOW_ROOT_DIR) --build-arg WOW_CLIENT_CONFIG=$(WOW_CLIENT_CONFIG) -f $(WOW_ROOT_DIR)/deployment/client/Dockerfile .
	docker run -d --name $(CLIENT_IMAGE) --network wow-network $(CLIENT_IMAGE)
	docker logs $(CLIENT_IMAGE)

# Stop and remove the server container
stop-server:
	docker stop $(SERVER_IMAGE) 2>/dev/null || true
	docker rm $(SERVER_IMAGE) 2>/dev/null || true

# Stop and remove the client container
stop-client:
	docker stop $(CLIENT_IMAGE) 2>/dev/null || true
	docker rm $(CLIENT_IMAGE) 2>/dev/null || true
