.PHONY: test test-report up

test:
	go test ./...

test-report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

###
IMAGE_NAME=packager
CONTAINER_NAME=packager
PORT=8080

up:
	docker build -t $(IMAGE_NAME) . && \
	docker stop $(CONTAINER_NAME) 2>/dev/null || true && \
	docker rm $(CONTAINER_NAME) 2>/dev/null || true && \
	docker run -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)