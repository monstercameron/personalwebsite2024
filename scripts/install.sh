#!/bin/bash

# scripts/install.sh

#!/bin/bash

echo "   /$$$$$$$                                                             /$$              " 
echo "  | $$__  $$                                                           | $$              " 
echo "  | $$  \ $$ /$$$$$$   /$$$$$$   /$$$$$$$  /$$$$$$  /$$$$$$$   /$$$$$$ | $$              " 
echo "  | $$$$$$$//$$__  $$ /$$__  $$ /$$_____/ /$$__  $$| $$__  $$ |____  $$| $$              " 
echo "  | $$____/| $$$$$$$$| $$  \__/|  $$$$$$ | $$  \ $$| $$  \ $$  /$$$$$$$| $$              " 
echo "  | $$     | $$_____/| $$       \____  $$| $$  | $$| $$  | $$ /$$__  $$| $$              " 
echo "  | $$     |  $$$$$$$| $$       /$$$$$$$/|  $$$$$$/| $$  | $$|  $$$$$$$| $$              " 
echo "  |__/      \_______/|__/      |_______/  \______/ |__/  |__/ \_______/|__/              " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "         /$$      /$$           /$$                 /$$   /$$                            " 
echo "        | $$  /$ | $$          | $$                |__/  | $$                            " 
echo "        | $$ /$$$| $$  /$$$$$$ | $$$$$$$   /$$$$$$$ /$$ /$$$$$$    /$$$$$$               " 
echo "        | $$/$$ $$ $$ /$$__  $$| $$__  $$ /$$_____/| $$|_  $$_/   /$$__  $$              " 
echo "        | $$$$_  $$$$| $$$$$$$$| $$  \ $$|  $$$$$$ | $$  | $$    | $$$$$$$$              " 
echo "        | $$$/ \  $$$| $$_____/| $$  | $$ \____  $$| $$  | $$ /$$| $$_____/              " 
echo "        | $$/   \  $$|  $$$$$$$| $$$$$$$/ /$$$$$$$/| $$  |  $$$$/|  $$$$$$$              " 
echo "        |__/     \__/ \_______/|_______/ |_______/ |__/   \___/   \_______/              " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "    /$$$$$$   /$$$$$$   /$$$$$$  /$$   /$$                                               " 
echo "   /$$__  $$ /$$$_  $$ /$$__  $$| $$  | $$                                               " 
echo "  |__/  \ $$| $$$$\ $$|__/  \ $$| $$  | $$                                               " 
echo "    /$$$$$$/| $$ $$ $$  /$$$$$$/| $$$$$$$$                                               " 
echo "   /$$____/ | $$\ $$$$ /$$____/ |_____  $$                                               " 
echo "  | $$      | $$ \ $$$| $$            | $$                                               " 
echo "  | $$$$$$$$|  $$$$$$/| $$$$$$$$      | $$                                               " 
echo "  |________/ \______/ |________/      |__/                                               " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "                                                                                         " 
echo "    /$$$$$$                      /$$           /$$   /$$ /$$$$$$$$ /$$      /$$ /$$   /$$" 
echo "   /$$__  $$                    | $$          | $$  | $$|__  $$__/| $$$    /$$$| $$  / $$" 
echo "  | $$  \__//$$$$$$   /$$$$$$  /$$$$$$        | $$  | $$   | $$   | $$$$  /$$$$|  $$/ $$/" 
echo "  | $$$$   /$$__  $$ |____  $$|_  $$_/        | $$$$$$$$   | $$   | $$ $$/$$ $$ \  $$$$/ " 
echo "  | $$_/  | $$$$$$$$  /$$$$$$$  | $$          | $$__  $$   | $$   | $$  $$$| $$  >$$  $$ " 
echo "  | $$    | $$_____/ /$$__  $$  | $$ /$$      | $$  | $$   | $$   | $$\  $ | $$ /$$/\  $$" 
echo "  | $$    |  $$$$$$$|  $$$$$$$  |  $$$$/      | $$  | $$   | $$   | $$ \/  | $$| $$  \ $$" 
echo "  |__/     \_______/ \_______/   \___/        |__/  |__/   |__/   |__/     |__/|__/  |__/" 
echo "                                                                                         " 
echo "                                                                                         " 
echo "                                                                                         " 


# Exit immediately if a command exits with a non-zero status
set -e

# Ensure old versions of Go are uninstalled and the new version is installed
required_version="go1.22.1"
go_tarball="go1.22.1.linux-amd64.tar.gz"

# Check if Go is installed and its version
echo "--> Checking Go installation..."
if command -v go &> /dev/null; then
    installed_version=$(go version | awk '{print $3}')
    if [ "$installed_version" != "$required_version" ]; then
        echo "|----> Go version $installed_version is installed. Removing it and installing Go $required_version..."
        sudo apt-get remove -y golang-go
        sudo apt-get remove -y --auto-remove golang-go
        wget https://go.dev/dl/$go_tarball
        sudo rm -rf /usr/local/go
        sudo tar -C /usr/local -xzf $go_tarball
        rm $go_tarball
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
        source ~/.profile
        echo "|----> Verifying Go installation..."
        go version
    else
        echo "|----> Go $required_version is already installed."
    fi
else
    echo "|----> Go is not installed. Installing Go $required_version..."
    wget https://go.dev/dl/$go_tarball
    sudo tar -C /usr/local -xzf $go_tarball
    rm $go_tarball
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    source ~/.profile
    echo "|----> Verifying Go installation..."
    go version
fi

echo "\n--> Setting up Go environment variables..."
echo 'export GOPATH=$HOME/go' >> ~/.profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.profile
source ~/.profile
echo "|----> GOPATH set to $HOME/go"
echo "|----> PATH updated with $GOPATH/bin"

echo "\n\n--> Starting installation of Personal Website 2024..."
echo "------------------------------------------------------------------"

# Step 1: Update system and install dependencies
echo -e "\n\n--> Step 1/10: Updating system and installing dependencies"
echo "------------------------------------------------------------------"
sudo apt update
echo "|----> System update complete"

echo "\n--> Checking npm and nodejs installation..."
# Check and install npm and nodejs
if ! command -v npm &> /dev/null; then
    echo "|----> npm is not installed. Installing npm and nodejs..."
    sudo apt install -y npm nodejs
else
    echo "|----> npm and nodejs are already installed."
fi

echo "\n--> Checking Git installation..."
# Check and install git
if ! command -v git &> /dev/null; then
    echo "|----> Git is not installed. Installing Git..."
    sudo apt install -y git
else
    echo "|----> Git is already installed."
fi

echo -e "\n\n--> System update and dependency installation complete."

# Step 2: Install templ
echo -e "\n\n--> Step 2/10: Installing templ"
echo "------------------------------------------------------------------"
go install github.com/a-h/templ/cmd/templ@latest
echo "|----> templ installed"

# Step 3: Set up project directory
echo -e "\n\n--> Step 3/10: Setting up project directory"
echo "------------------------------------------------------------------"
PROJECT_DIR="$HOME/PersonalWebsite2024"
if [ -d "$PROJECT_DIR" ]; then
    echo "|----> Project directory already exists. Removing..."
    rm -rf $PROJECT_DIR
fi
mkdir -p $PROJECT_DIR
cd $PROJECT_DIR
echo "|----> Done setting up project directory"

# Step 4: Clone the repository
echo -e "\n\n--> Step 4/10: Cloning the repository"
echo "------------------------------------------------------------------"
if [ -d ".git" ]; then
    echo "|----> Existing repository found. Removing..."
    rm -rf .git
fi
git clone https://github.com/monstercameron/PersonalWebsite2024.git .
echo "|----> Repository cloned"

# Step 5: Install Go dependencies
echo -e "\n\n--> Step 5/10: Installing Go dependencies"
echo "------------------------------------------------------------------"
go mod download
echo "|----> Go dependencies installed"

# Step 6: Install npm dependencies
echo -e "\n\n--> Step 6/10: Installing npm dependencies"
echo "------------------------------------------------------------------"
npm install
echo "|----> npm dependencies installed"

# Step 7: Generate templ files
echo -e "\n\n--> Step 7/10: Generating templ files"
echo "------------------------------------------------------------------"
cd views
templ generate
cd ..
echo "|----> templ files generated"

# Step 8: Build Tailwind CSS
echo -e "\n\n--> Step 8/10: Building Tailwind CSS"
echo "------------------------------------------------------------------"
npx tailwindcss -i ./styles/input.css -o ./static/output.css
echo "|----> Tailwind CSS built"

# Step 9: Create .env file
echo -e "\n\n--> Step 9/10: Creating .env file"
echo "------------------------------------------------------------------"
env_file=".env"
if [ -f "$env_file" ]; then
    echo "|----> .env file already exists. Removing..."
    rm $env_file
fi
cat > .env << EOL
PORT=80
OAIKEY=your_openai_api_key_here
VSCODEUSER=your_vscode_username
VSCODEPASS=your_vscode_password
EOL
echo "|----> .env file created"

# Step 10: Build the Go application
echo -e "\n\n--> Step 10/10: Building the Go application"
echo "------------------------------------------------------------------"
go build -o PersonalWebsite2024 main.go
echo "|----> Go application built"

# Set up systemd service
echo -e "\n\n--> Setting up systemd service..."
echo "------------------------------------------------------------------"
service_file="/etc/systemd/system/personalwebsite.service"
if [ -f "$service_file" ]; then
    echo "|----> systemd service file already exists. Removing..."
    sudo rm $service_file
fi
sudo bash -c 'cat > /etc/systemd/system/personalwebsite.service << EOL
[Unit]
Description=Personal Website Service
After=network.target

[Service]
ExecStart=$HOME/PersonalWebsite2024/PersonalWebsite2024
WorkingDirectory=$HOME/PersonalWebsite2024
User=$(whoami)
Restart=always
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=personalwebsite

[Install]
WantedBy=multi-user.target
EOL'
echo "|----> systemd service file created"

# Reload systemd, enable and start the service
sudo systemctl daemon-reload
sudo systemctl enable personalwebsite.service
sudo systemctl start personalwebsite.service
echo "|----> systemd service started"

echo -e "\n\n--> Installation complete!"
echo "------------------------------------------------------------------"
echo "|----> Please edit the .env file with your actual values:"
echo "|----> nano $HOME/PersonalWebsite2024/.env"
echo "|----> The personal website service is now running and will start automatically on boot."
echo "|----> You can check its status with: sudo systemctl status personalwebsite.service"
