# Neverpay API User Service Project
The code structure and conventions inspired by tutorialedge.net

## Requirements to run the application in development environment
- Docker (required)
- VSCode with devContainer extension installed (optional, recommended)

## Run the application in the development environment
1. Optional - If you wish to use VSCode devContainer, install [the plugin first](https://code.visualstudio.com/docs/remote/create-dev-container)
2. Optional - Re-open the project inside a devcontainer (if you open the project in VSCode with devcontainer plugin installed, you should see a pop-up asking you to open the project inside a devcontainer)
3. Create a .env file at the root level of the project with the following environment variables.
```
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_DB=mypostgresdb
POSTGRES_USER=mypostgresuser
POSTGRES_PASSWORD=mypostgrespassword
SSL_MODE=disable
AUTH_SECRET=yourauthsecretkey
```
4. You should be able to run the server with the following command
``` 
docker-compose up --build
```
