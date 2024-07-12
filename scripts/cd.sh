#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Change to the PersonalWebsite2024 directory
cd /root/PersonalWebsite2024

# Pull the latest changes from the main branch
git restore .
git pull

# Generate templ files
cd views
/root/go/bin/templ generate
cd ..

# Build Tailwind CSS
npx tailwindcss -i ./tailwinds/input.css -o ./static/output.css

# Rebuild the Go application
/usr/local/go/bin/go build -o PersonalWebsite2024

# Restart the PersonalWebsite service
systemctl restart personalwebsite

echo "PersonalWebsite2024 updated and restarted successfully"