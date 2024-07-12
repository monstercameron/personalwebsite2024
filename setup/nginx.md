# Basic Installation Guide for Personal Website and Nginx on Ubuntu

This guide provides detailed steps for setting up your personal website as a systemd service and configuring Nginx as a reverse proxy on an Ubuntu system.

## Prerequisites

- Ubuntu operating system
- Root or sudo privileges
- Nginx installed

## Step 1: Setting Up the Nginx Configuration

1. **Open the Nginx Configuration File**:
   ```bash
   sudo nano /etc/nginx/sites-available/personalwebsite
   ```

   ### Explanation:
   This command opens the `personalwebsite` configuration file in the `nano` text editor. The file is located in the `/etc/nginx/sites-available/` directory, which is where individual site configurations for Nginx are typically stored. If the file does not exist, this command will create it.

2. **Configure Nginx**: Add the following configuration to set up a reverse proxy for your personal website running on port 8080:

   ```nginx
   # Nginx server block for www.earlcameron.com
   server {
       server_name www.earlcameron.com;
       listen 443 ssl;

       # Define the root for static assets
       root /root/PersonalWebsite2024/static;

       # SSL configurations managed by Certbot
       ssl_certificate /etc/letsencrypt/live/www.earlcameron.com/fullchain.pem;
       ssl_certificate_key /etc/letsencrypt/live/www.earlcameron.com/privkey.pem;
       include /etc/letsencrypt/options-ssl-nginx.conf;
       ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

       # Proxy requests to the Go application
       location / {
           proxy_pass http://127.0.0.1:8080;  # Update to the port where your Go app is running
           proxy_http_version 1.1;
           proxy_set_header Upgrade $http_upgrade;
           proxy_set_header Connection 'upgrade';
           proxy_set_header Host $host;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
           proxy_cache_bypass $http_upgrade;
           proxy_read_timeout 60s;
       }
   }

   # Redirect HTTP requests to HTTPS
   server {
       listen 80;
       server_name www.earlcameron.com;

       location / {
           return 301 https://$host$request_uri;
       }
   }

   # Redirect non-www to www and HTTP to HTTPS
   server {
       listen 80;
       server_name earlcameron.com;

       location / {
           return 301 https://www.earlcameron.com$request_uri;
       }
   }

   server {
       listen 443 ssl;
       server_name earlcameron.com;

       ssl_certificate /etc/letsencrypt/live/www.earlcameron.com/fullchain.pem;
       ssl_certificate_key /etc/letsencrypt/live/www.earlcameron.com/privkey.pem;
       include /etc/letsencrypt/options-ssl-nginx.conf;
       ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

       location / {
           return 301 https://www.earlcameron.com$request_uri;
       }
   }

   # Nginx server block for reader.earlcameron.com
   server {
       server_name reader.earlcameron.com;
       listen 80;

       location / {
           proxy_pass http://127.0.0.1:3001/;
           proxy_set_header Host $host;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       }
   }
   ```

   ### Explanation:
   This configuration defines multiple server blocks:
   - The first block listens for HTTPS requests to `www.earlcameron.com`, serves static files from `/root/PersonalWebsite2024/static`, and proxies requests to your Go application running on port 8080.
   - The second block redirects HTTP requests to `www.earlcameron.com` to HTTPS.
   - The third and fourth blocks handle redirects from `earlcameron.com` to `www.earlcameron.com` and ensure that all traffic is served over HTTPS.
   - The final block proxies requests to `reader.earlcameron.com` to another application running on port 3001.

3. **Enable the Nginx Configuration**:
   ```bash
   sudo ln -s /etc/nginx/sites-available/personalwebsite /etc/nginx/sites-enabled/
   ```

   ### Explanation:
   This command creates a symbolic link from the configuration file in `sites-available` to `sites-enabled`, which is where Nginx reads the active configuration files.

4. **Test the Nginx Configuration**:
   ```bash
   sudo nginx -t
   ```

   ### Explanation:
   This command tests the Nginx configuration for syntax errors. If there are any issues, it will display an error message.

5. **Restart Nginx**:
   ```bash
   sudo systemctl restart nginx
   ```

   ### Explanation:
   This command restarts the Nginx service to apply the new configuration.

## Step 2: Check Nginx Status

- To check if Nginx is running:
  ```bash
  sudo systemctl status nginx
  ```

  ### Explanation:
  This command checks the status of the Nginx service, showing whether it is active (running) or inactive (stopped), along with other details.

## Step 3: Setting Up the Personal Website as a Systemd Service

1. **Create the Systemd Service File**:
   ```bash
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
   ```

   ### Explanation:
   This script creates a systemd service file for the personal website. It defines the service name, the command to start the service, the working directory, the user to run the service, and how to handle service restarts.

2. **Reload Systemd, Enable and Start the Service**:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable personalwebsite.service
   sudo systemctl start personalwebsite.service
   echo "|----> systemd service started"
   ```

   ### Explanation:
   These commands reload the systemd manager configuration, enable the personal website service to start on boot, and start the service immediately.

3. **Installation Complete**:
   ```bash
   echo -e "\n\n--> Installation complete!"
   echo "------------------------------------------------------------------"
   echo "|----> Please edit the .env file with your actual values:"
   echo "|----> nano $HOME/PersonalWebsite2024/.env"
   echo "|----> The personal website service is now running and will start automatically on boot."
   echo "|----> You can check its status with: sudo systemctl status personalwebsite.service"
   ```

   ### Explanation:
   This message confirms that the installation is complete and instructs the user to edit the `.env` file with actual values. It also provides a reminder to check the service status.

## Step 4: Interacting with the Personal Website Service

- **Check the status**:
  ```bash
  sudo systemctl status personalwebsite
  ```

- **Start the service**:
  ```bash
  sudo systemctl start personalwebsite
  ```

- **Stop the service**:
  ```bash
  sudo systemctl stop personalwebsite
  ```

- **Restart the service**:
  ```bash
  sudo systemctl restart personalwebsite
  ```

- **Enable the service to start on boot**:
  ```bash
  sudo systemctl enable personalwebsite
  ```

- **Disable the service from starting on boot**:
  ```bash
  sudo systemctl disable personalwebsite
  ```

### Explanation:
These commands allow you to manage the personal website service, including checking its status, starting, stopping, restarting, enabling, and disabling it.

## Notes

- The Nginx configuration file is located at: `/etc/nginx/sites-available/personalwebsite`
- The script for creating the systemd service file is part of the installation script located at: `scripts/install.sh`
- Ensure that the paths and user settings in the service file and Nginx configuration are correct and match your system's setup.

By following these steps, you will have your personal website set up as a systemd service and Nginx configured to reverse proxy it, ensuring your site is accessible and managed correctly.