
QUERYING_BINARY=queryinglogs
Q_WRITER_BINARY=queueWriter
DB_WRITER_BINARY=dbWriter


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up 
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build:
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build 
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_querying: builds the querying binary as a linux executable
build_querying:
	@echo "Building querying-logs binary..."
	cd ../querying-logs && env GOOS=linux CGO_ENABLED=0 go build -o ${QUERYING_BINARY} .
	@echo "Done!"

## build_queue-writer: builds the queue-writer binary as a linux executable
build_queue-writer:
	@echo "Building queue-writer binary..."
	cd ../queue-writer && env GOOS=linux CGO_ENABLED=0 go build -o ${Q_WRITER_BINARY} .
	@echo "Done!"

## build_db-writer: builds the db-writer binary as a linux executable
build_db-writer:
	@echo "Building db-writer binary..."
	cd ../db-writer && env GOOS=linux CGO_ENABLED=0 go build -o ${DB_WRITER_BINARY} .
	@echo "Done!"



