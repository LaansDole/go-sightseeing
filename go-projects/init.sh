#!/bin/bash

# Check if BASE_DIR argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <base_directory>"
  exit 1
fi

BASE_DIR=$1

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
echo "Project structure for Banking Service has been set up in ${BASE_DIR}."

# Open the project directory
cd "${BASE_DIR}" || echo -e "Cannot find directory" && return
go mod init "${BASE_DIR}"