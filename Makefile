build:
	@if [ -f bin/app ]; then \
		rm bin/app; \
		echo "Deleted bin/app"; \
	fi
	@echo "Building binary..."
	@go build -o bin/app

concurency: build
	@./bin/app concurency

test:
	go test -v ./...