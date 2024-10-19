#!/bin/bash

if command -v go &> /dev/null; then
	echo "Go is installed."
	echo "Generating TLS certification."
else
	echo "Go doesn't seem to be installed on your system."
	echo "Please install Go and then try again."
fi

mkdir tls
cd tls
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host localhost

echo "TLS certificates written to tls/ directory."
