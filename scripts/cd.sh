#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Change to the PersonalWebsite2024 directory
cd /root/PersonalWebsite2024

# Pull the latest changes from the main branch
git pull origin main

# Generate templ files
cd views
templ generate
cd ..

# Build Tailwind CSS
npx tailwindcss -i ./styles/input.css -o ./static/output.css

# Rebuild the Go application
go build -o PersonalWebsite2024 main.go

# Restart the PersonalWebsite service
systemctl restart personalwebsite

echo "PersonalWebsite2024 updated and restarted successfully"