#!/bin/bash

# Function to display help
display_help() {
    echo "Usage: init_setup.sh -d <base_directory>"
    echo
    echo "Options:"
    echo "  -d, --directory    Specify the base directory to set up the project structure"
    echo "  -h, --help         Display this help message"
    return 0
}

# Check if no arguments are provided
if [ "$#" -eq 0 ]; then
    display_help
    return 0
fi

# Parse the command-line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -d|--directory)
            BASE_DIR="$2"
            shift 2
            ;;
        -h|--help)
            display_help
            ;;
        *)
            echo "Unknown option: $1"
            display_help
            ;;
    esac
done

# Check if BASE_DIR is set
if [ -z "$BASE_DIR" ]; then
    echo "Error: Base directory not specified."
    display_help
fi

# Create directories
mkdir -p "${BASE_DIR}/cmd/api"
mkdir -p "${BASE_DIR}/docs"
mkdir -p "${BASE_DIR}/internal/api/adapter/in"
mkdir -p "${BASE_DIR}/internal/api/adapter/out"
mkdir -p "${BASE_DIR}/internal/api/port/in"
mkdir -p "${BASE_DIR}/internal/api/port/out"
mkdir -p "${BASE_DIR}/internal/api/service"
mkdir -p "${BASE_DIR}/internal/common/connect/fabric"
mkdir -p "${BASE_DIR}/internal/common/connect/temporal"
mkdir -p "${BASE_DIR}/internal/common/domain/stub"
mkdir -p "${BASE_DIR}/internal/worker/adapter/out"
mkdir -p "${BASE_DIR}/internal/worker/port/out"
mkdir -p "${BASE_DIR}/internal/worker/service"

# Create a basic main.go file
cat <<EOF > "${BASE_DIR}/cmd/api/main.go"
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting Banking Service API...")
	// Initialize gRPC server and Temporal client here
	log.Fatal("Implement gRPC server and Temporal client initialization")
}
EOF

# Print success message
echo "Project structure has been set up in ${BASE_DIR}."

# Open the project directory and initialize go module
# shellcheck disable=SC2015
cd "${BASE_DIR}" && go mod init "$(basename "${BASE_DIR}")" || return 0
