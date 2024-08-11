BINARY_NAME=main

all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME)

# Run the Go binary
run: build
	./$(BINARY_NAME)

# Clean up the build artifacts
clean:
	rm -f $(BINARY_NAME)