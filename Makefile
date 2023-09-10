build:
	@if [ -f bin/app ]; then \
		rm bin/app; \
		echo "Found an existing binary..."; \
		echo "Deleted existing bin/app binary"; \
	fi
	@echo "Building a new binary in the bin folder..."
	@go build -o bin/app

concurency: build
	@./bin/app concurency

userprofile: build
	@./bin/app userprofile

test:
	go test -v ./...