#!/bin/bash

# Check if swag is installed
if ! command -v swag &> /dev/null
then
    echo "swag could not be found, installing..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Navigate to the project root directory where main.go is located (if necessary)
# cd /path/to/your/project-directory

# Run swag init with the -g flag to specify the location of the main API file
swag init -g cmd/main.go