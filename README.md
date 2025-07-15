# Calorie Tracker

An application that helps you keep track of your calorie goals. Written with ❤️ using Golang and Vue.js.

## Installation
1. Setup a user for the calorietracking application
2. Run the table setup script in /migrations
3. Grant all permissions on the table to the user.
4. Make a copy of .env.example to .env, and fill out the variables as needed.
    - Example values are provided in the .env.example.

### With Docker Compose
```yaml
services:
    calorierouter:
        image: calorierouter:latest
        ports:
            - "8080:8080"
        env_file: .env
```

## Build from source
1. Clone the repository
2. Using the tailwind CLI, run `tailwindcss -i /app/public/css/input.css -o /app/public/css/index.css`
3. In the /app/router directory, run `go build .`

## Screenshots


