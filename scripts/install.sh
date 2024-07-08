#!/bin/bash

# scripts/install.sh

# Exit immediately if a command exits with a non-zero status
set -e

# Function to display a simple progress bar
progress_bar() {
    local duration=$1
    local steps=$2
    local step_duration=$(echo "scale=2; $duration / $steps" | bc)
    local progress=0
    while [ $progress -le $steps ]; do
        echo -ne "\r[${progress}/${steps}] ["
        for ((i=0; i<progress; i++)); do echo -n "="; done
        for ((i=progress; i<steps; i++)); do echo -n " "; done
        echo -n "]"
        sleep $step_duration
        ((progress++))
    done
    echo
}

echo "Starting installation of Personal Website 2024..."

# Step 1: Update system and install dependencies
echo "Step 1/10: Updating system and installing dependencies"
sudo apt-get update
sudo apt-get install -y golang-go npm
progress_bar 5 10

# Step 2: Install templ
echo "Step 2/10: Installing templ"
go install github.com/a-h/templ/cmd/templ@latest
progress_bar 3 10

# Step 3: Set up project directory
echo "Step 3/10: Setting up project directory"
PROJECT_DIR="/root/PersonalWebsite2024"
mkdir -p $PROJECT_DIR
cd $PROJECT_DIR
progress_bar 1 10

# Step 4: Clone the repository
echo "Step 4/10: Cloning the repository"
git clone https://github.com/monstercameron/PersonalWebsite2024.git .
progress_bar 5 10

# Step 5: Install Go dependencies
echo "Step 5/10: Installing Go dependencies"
go mod download
progress_bar 3 10

# Step 6: Install npm dependencies
echo "Step 6/10: Installing npm dependencies"
npm install
progress_bar 5 10

# Step 7: Generate templ files
echo "Step 7/10: Generating templ files"
cd views
templ generate
cd ..
progress_bar 2 10

# Step 8: Build Tailwind CSS
echo "Step 8/10: Building Tailwind CSS"
npx tailwindcss -i ./styles/input.css -o ./static/output.css
progress_bar 3 10

# Step 9: Create .env file
echo "Step 9/10: Creating .env file"
cat > .env << EOL
PORT=80
OAIKEY=your_openai_api_key_here
VSCODEUSER=your_vscode_username
VSCODEPASS=your_vscode_password
EOL
progress_bar 1 10

# Step 10: Build the Go application
echo "Step 10/10: Building the Go application"
go build -o PersonalWebsite2024 main.go
progress_bar 3 10

# Set up systemd service
echo "Setting up systemd service..."
cat > /etc/systemd/system/personalwebsite.service << EOL
[Unit]
Description=Personal Website Service
After=network.target

[Service]
ExecStart=/root/PersonalWebsite2024/PersonalWebsite2024
WorkingDirectory=/root/PersonalWebsite2024
User=root
Restart=always
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=personalwebsite

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd, enable and start the service
sudo systemctl daemon-reload
sudo systemctl enable personalwebsite.service
sudo systemctl start personalwebsite.service

echo "Installation complete!"
echo "Please edit the .env file with your actual values:"
echo "nano /root/PersonalWebsite2024/.env"
echo "The personal website service is now running and will start automatically on boot."
echo "You can check its status with: sudo systemctl status personalwebsite.service"