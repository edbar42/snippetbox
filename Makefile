.PHONY: all run make_cert clean

# Build the project
all: build

build:
	go build -o snippetbox cmd/web/*

# Run the project
run: build
	./snippetbox

# Run test suite
test:
	go test -v ./cmd/web ./pkg/models/mysql/

# Generate TLS certificates
tls:
	@echo "Checking if Go is installed..."
	if command -v go &> /dev/null; then \
		echo "Go is installed."; \
		echo "Generating TLS certification."; \
		mkdir -p tls && cd tls && go run $$(go env GOROOT)/src/crypto/tls/generate_cert.go --host localhost; \
		echo "TLS certificates written to tls/ directory."; \
	else \
		echo "Go doesn't seem to be installed on your system."; \
		echo "Please install Go and then try again."; \
	fi

# Clean up the build
clean:
	rm -f snippetbox
	rm -rf tls

