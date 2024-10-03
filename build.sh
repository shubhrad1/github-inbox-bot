#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Download all dependencies
echo "Downloading dependencies..."
go mod download

# Load environment variables
echo "Loading environment variables..."
touch .env
echo "GITHUB_TOKEN=${GITHUB_TOKEN}" > .env
echo "SLACK_TOKEN=${SLACK_TOKEN}" > .env
echo "SLACK_CHANNEL=${SLACK_CHANNEL}" > .env
echo "GHTOKEN_START_DATE=${GHTOKEN_START_DATE}" > .env


# Source the .env file to load the environment variables
source .env

# Build the Go app
echo "Building the Go app..."
go build -o github-inbox-bot

echo "Build completed successfully."