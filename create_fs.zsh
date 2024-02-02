#!/bin/zsh

# Define the project root directory
PROJECT_ROOT="."

# Create the directory structure
mkdir -p "$PROJECT_ROOT/cmd/app"
mkdir -p "$PROJECT_ROOT/internal/handler"
mkdir -p "$PROJECT_ROOT/internal/model"
mkdir -p "$PROJECT_ROOT/internal/repository"
mkdir -p "$PROJECT_ROOT/pkg/config"
mkdir -p "$PROJECT_ROOT/pkg/database"

# Create the necessary files
touch "$PROJECT_ROOT/cmd/app/main.go"
touch "$PROJECT_ROOT/internal/handler/message.go"
touch "$PROJECT_ROOT/internal/model/message.go"
touch "$PROJECT_ROOT/internal/repository/message.go"
touch "$PROJECT_ROOT/pkg/config/config.go"
touch "$PROJECT_ROOT/pkg/database/database.go"
touch "$PROJECT_ROOT/.env"
touch "$PROJECT_ROOT/Dockerfile"
touch "$PROJECT_ROOT/go.mod"
touch "$PROJECT_ROOT/go.sum"
touch "$PROJECT_ROOT/README.md"

# Print a success message
echo "Project directories and files created successfully in $PROJECT_ROOT"
