# Earl Cameron's Personal Website

## Overview

This project is a personal website for Earl Cameron, a Full-stack Developer and AI Innovator. The website showcases Earl's professional experience, skills, and projects, and includes integration with OpenAI's GPT and DALL-E models for text and image generation.

## Technology Stack

- **Backend:** Go (Golang)
- **Frontend:** HTML, CSS (Tailwind CSS), JavaScript (HTMX)
- **Templating:** templ
- **AI Integration:** OpenAI API (GPT-3.5 and DALL-E)
- **Environment Management:** godotenv

## Project Structure

```
personalwebsite/
├── main.go
├── openai/
│   └── openai_handler.go
├── views/
│   ├── home.templ
│   ├── resume.templ
│   ├── layout.templ
│   └── nav.templ
├── static/
│   └── output.css
├── styles/
│   └── input.css
├── .env
├── go.mod
├── go.sum
└── README.md
```

## Features

- Responsive design using Tailwind CSS
- Dynamic content loading with HTMX
- Integration with OpenAI for text and image generation
- Resume page showcasing professional experience and skills
- Modular template structure using templ

## Setup and Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/personalwebsite.git
   cd personalwebsite
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   npm install
   ```

3. Set up environment variables:
   Create a `.env` file in the root directory with the following content:
   ```
   OAIKEY=your_openai_api_key_here
   PORT=8080
   ```

4. Generate CSS:
   ```sh
   npx tailwindcss -i ./styles/input.css -o ./static/output.css
   ```

5. Generate templ files:
   ```sh
   templ generate
   ```

6. Run the server:
   ```sh
   go run main.go
   ```
   The server will start on [http://localhost:8080](http://localhost:8080) (or the port specified in your `.env` file).

## Development

- To watch for changes in Tailwind CSS:
  ```sh
  npx tailwindcss -i ./styles/input.css -o ./static/output.css --watch
  ```

- To regenerate templ files after changes:
  ```sh
  templ generate
  ```

- To run the server with hot reloading, you can use tools like `air`:
  ```sh
  air
  ```

## API Endpoints

- `/`: Home page
- `/resume`: Resume page
- `/generate-text?prompt=<your_prompt>`: Generate text using OpenAI's GPT model
- `/generate-image?prompt=<your_prompt>`: Generate an image using DALL-E

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Contact

Earl Cameron - [mr.e.cameron@gmail.com](mailto:mr.e.cameron@gmail.com)

Project Link: [https://github.com/yourusername/personalwebsite](https://github.com/yourusername/personalwebsite)

## Acknowledgements

- [Go](https://golang.org/)
- [HTMX](https://htmx.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [templ](https://github.com/a-h/templ)
- [OpenAI](https://openai.com/)
- [godotenv](https://github.com/joho/godotenv)
```