# Telegram bot (golang)

A basic example of a telegram bot answering questions with commands.
It searches in wikipedia about different subjects and gives a brief
information about it. Results are cached in JSON format to keep requests
limited.

1. Create environment variables

   ```sh
   $ cp .env.example .env
   ```

2. Install dependencies

   ```sh
   $ go mod download
   ```

3. Run the application
   ```sh
   $ go run cmd/main.go
   ```
