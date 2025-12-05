.PHONY: all build-web build-go run

# Default target: builds everything
all: build-web build-go

# 1. Build SvelteKit & Copy to internal/ui
build-web:
	@echo "Building Web UI..."
	cd web && npm install && npm run build
	@echo "Copying web assets to internal/ui..."
	rm -rf internal/ui/build
	cp -r web/build internal/ui/build

# 2. Build Go Binary
build-go:
	@echo "Building Go Binary..."
	go build -o bin/core .

# 3. Run the full app
run: build-web build-go
	@echo "Starting App..."
	./bin/core serve