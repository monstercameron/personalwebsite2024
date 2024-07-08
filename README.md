# Personal Website 2024

A cutting-edge personal website showcasing modern web development techniques and AI integration.

## Project Overview

- **Backend:** Go (Golang) 1.22.1
- **Frontend:** HTML5 with HTMX 2.0.0 for dynamic content loading
- **Styling:** Tailwind CSS 3.4.1
- **Templating:** templ 0.2.598
- **AI Integration:** OpenAI API (GPT-3.5 and DALL-E models)
- **Database:** JSON files for content storage
- **Deployment:** Optimized for Digital Ocean VPS

## Key Features

1. **Responsive Design:**
   - Mobile-first approach using Tailwind CSS
   - Seamless experience across devices (mobile, tablet, desktop)

2. **Dynamic Content Loading:**
   - HTMX integration for SPA-like experience without heavy JavaScript
   - Smooth transitions between pages

3. **AI-Powered Functionalities:**
   - Text generation using GPT-3.5 model
   - Image generation with DALL-E
   - Daily quote generation with caching mechanism

4. **Content Management:**
   - Blog system with JSON-based content storage
   - Easy-to-update project showcase
   - Dynamic resume page

5. **AI Workshop:**
   - Secure login system
   - Integration with remote Visual Studio Code
   - Playground for AI experiments

6. **Performance Optimizations:**
   - Efficient routing system
   - Static file serving with proper caching
   - Minimal JavaScript usage

7. **Monitoring and Logging:**
   - Custom middleware for request logging
   - Metrics endpoint for basic analytics

## Project Structure

```
/
├── main.go           # Main application entry point
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── .env              # Environment variables (not in repo)
├── static/           # Static files (CSS, images, etc.)
│   └── output.css    # Compiled Tailwind CSS
├── views/            # templ template files
│   ├── layout.templ
│   ├── home.templ
│   ├── blog.templ
│   └── ... 
├── styles/
│   └── input.css     # Tailwind CSS input file
├── routes/
│   └── router.go     # Route definitions and handlers
├── middleware/
│   └── logger.go     # Logging middleware
├── openai/
│   └── openai_handler.go  # OpenAI API integration
├── utils/
│   └── blog_utils.go # Utility functions for blog
└── scripts/
    └── install.sh    # Installation script
```

## Installation

### Prerequisites

- Linux-based system (Ubuntu 22.04 LTS recommended)
- Root access or sudo privileges
- Git installed
- Internet connection for downloading dependencies

### Detailed Installation Steps

1. **SSH into your server:**
   ```bash
   ssh root@your_server_ip
   ```

2. **Create the installation script:**
   ```bash
   nano /root/install.sh
   ```

3. **Copy the contents of the `install.sh` script (provided separately) into this file. Save and exit (Ctrl+X, then Y, then Enter).**

4. **Make the script executable:**
   ```bash
   chmod +x /root/install.sh
   ```

5. **Run the installation script:**
   ```bash
   ./root/install.sh
   ```

6. **The script will:**
   - Update system packages
   - Install Go, npm, and other dependencies
   - Clone the repository
   - Set up the project structure
   - Generate `templ` files
   - Build Tailwind CSS
   - Create a systemd service for automatic startup

7. **After installation, edit the `.env` file with your actual values:**
   ```bash
   nano /root/PersonalWebsite2024/.env
   ```

   Required values:
   - `PORT`: The port your server will run on (default: 80)
   - `OAIKEY`: Your OpenAI API key
   - `VSCODEUSER`: Username for AI Workshop access
   - `VSCODEPASS`: Password for AI Workshop access

8. **Restart the service to apply the new environment variables:**
   ```bash
   sudo systemctl restart personalwebsite.service
   ```

## Usage Guide

1. **Accessing the Website:**
   - Navigate to `http://your_server_ip` in a web browser
   - If you've set up a domain, use that instead

2. **Navigation:**
   - Use the top navigation bar to explore different sections
   - Experience smooth transitions powered by HTMX

3. **Blog Management:**
   - Edit the `/root/PersonalWebsite2024/static/blog_posts.json` file to update blog content
   - Format:
     ```json
     {
       "posts": [
         {
           "title": "Post Title",
           "slug": "post-slug",
           "content": "Post content in HTML format",
           "date": "2024-03-15T00:00:00Z"
         },
         ...
       ]
     }
     ```

4. **AI Workshop:**
   - Navigate to the AI Workshop section
   - Log in using the credentials set in the `.env` file
   - Experiment with text and image generation

5. **Updating the Website:**

   For code changes:
   - SSH into the server
   - Navigate to `/root/PersonalWebsite2024`
   - Make your changes
   - Rebuild: `go build -o PersonalWebsite2024 main.go`
   - Restart: `sudo systemctl restart personalwebsite.service`

   For style changes:
   - Edit the `styles/input.css` file
   - Rebuild CSS: `npx tailwindcss -i ./styles/input.css -o ./static/output.css`

## Maintenance and Monitoring

1. **Check service status:**
   ```bash
   sudo systemctl status personalwebsite.service
   ```

2. **View logs:**
   ```bash
   journalctl -u personalwebsite.service
   ```

3. **Monitor system resources:**
   ```bash
   top or htop
   ```

4. **Check disk usage:**
   ```bash
   df -h
   ```

5. **View real-time log:**
   ```bash
   tail -f /var/log/syslog | grep personalwebsite
   ```

## Acknowledgements

- Go community for the excellent programming language and tools
- HTMX for simplifying dynamic content without heavy JavaScript
- Tailwind CSS for the utility-first CSS framework
- OpenAI for their powerful API
- Digital Ocean for reliable hosting solutions