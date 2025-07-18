# 🍪 Calorie Tracker 🍪

An application that helps you keep track of your calorie goals. Written with ❤️ using Golang and Vue.js.

<img width="1920" height="878" alt="The login screen for the Calorie Tracker application" src="https://github.com/user-attachments/assets/8b38ca0f-1e7a-46fc-b32c-8df4cb61e7bd" />

## Installation

This application requires a MariaDB instance, with a user and database created for this application. The user must have all permissions on the application's database.

### With Docker Compose
1. Run the migrations script in /migrations
2. Make a copy of .env.example to .env, and fill out the variables as needed.
3. Refer to the docker compose file below.
```yaml
services:
    calorierouter:
        image: ghcr.io/Inventhrice/calorierouter:latest
        ports:
            - "8080:8080"
        env_file: .env
        restart: unless_stopped
        depends_on: [name of mariadb service]
```

## Build from source
1. Clone the repository
2. Using the tailwind CLI, run `tailwindcss -i /app/public/css/input.css -o /app/public/css/index.css`
3. In the /app/router directory, run `go build .`

## Screenshots
<img width="1920" height="878" alt="image" src="https://github.com/user-attachments/assets/9c65e260-645c-443f-ba5d-ec9b91c447ed" />
<img width="1920" height="878" alt="image" src="https://github.com/user-attachments/assets/a83c6d96-abc9-45b4-b5db-417f46511da5" />





