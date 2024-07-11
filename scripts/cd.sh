#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Change to the PersonalWebsite2024 directory
cd /root/PersonalWebsite2024

# Pull the latest changes from the main branch
git pull

# Generate templ files
npm run templ
cd /root/PersonalWebsite2024

# Build Tailwind CSS
npm run css

# Rebuild the Go application
go build -o PersonalWebsite2024

# Restart the PersonalWebsite service
systemctl restart personalwebsite

echo "PersonalWebsite2024 updated and restarted successfully"