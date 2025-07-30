# 🍪 Calorie Tracker 🍪

An application that helps you keep track of your calorie goals. Written with ❤️ using Golang and Vue.js.

<img width="1920" height="878" alt="The login screen for the Calorie Tracker application" src="https://github.com/user-attachments/assets/8b38ca0f-1e7a-46fc-b32c-8df4cb61e7bd" />

## Installation

This application requires a MariaDB instance, with a user and database created for this application. The user must have all permissions on the application's database.

### With Docker Compose
1. Make a copy of .env.example to .env, and fill out the variables as needed.
2. Refer to the docker compose file below.
```yaml
services:
    mariadb:
        container_name: calorierouter_db
        image: lscr.io/linuxserver/mariadb
        environment:
        - PUID=1000
        - PGID=1000
        - TZ=EST
        - MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD}
        - MYSQL_DATABASE=${DB_DBNAME}
        - MYSQL_USER=${DB_USER}
        - MYSQL_PASSWORD=${DB_PASSWORD}
        volumes:
            - ./mariadb:/var/lib/mysql/
        ports:
            - "3307:3306"
    calorierouter:
        image: ghcr.io/Inventhrice/calorierouter:latest
        ports:
            - "8080:8080"
        env_file: .env
        restart: unless_stopped
        depends_on: 
            - mariadb
```
3. Start the database container using `docker compose up -d mariadb`
4. Create the inital database using the following command: 
5. Start the application using `docker compose up -d calorierouter`

## Building
### From the Dockerfile
Execute `docker build .` in the /app directory.

### From Source
**Requirements:** npm and golang must be installed on the system.
1. Clone the repository
2. In the /app/frontend_src directory, run `npm install`, then run `npm run build`. This step will automatically make a directory called public.
3. In the /app/router directory, run `go run .`

## Development Environment
### Frontend
This does not come with docker instructions, as I recommend you use npm on your local machine. Please follow the relavent instructions under the "Building > From Source" section.

### Backend
Execute `docker compose -f docker-compose.dev.yml up -d calrouter`.

### Complete
Execute `docker compose -f docker-compose.dev.yml build calorietracker` to compile a complete image.

## Screenshots
<img width="1920" height="878" alt="image" src="https://github.com/user-attachments/assets/9c65e260-645c-443f-ba5d-ec9b91c447ed" />
<img width="1920" height="878" alt="image" src="https://github.com/user-attachments/assets/a83c6d96-abc9-45b4-b5db-417f46511da5" />