# Neverpay API Project
The code structure and conventions inspired by tutorialedge.net

## Run the application in the development environment
1. If you wish to use VSCode devContainer, install [the plugin first](https://code.visualstudio.com/docs/remote/create-dev-container). Make sure you have Docker installed and running.
2. Re-open the project inside a devcontainer (if you open the project in VSCode with devcontainer plugin installed, you should see a pop-up asking you to open the project inside a devcontainer)
3. Create a .env file in the root level of the project with the following environment variables. NOTE: If you are using VSCode devContainer, and your postgres is running in another container in your machine, use ```host.docker.internal``` as your DB_HOST instad of ```localhost```
```
POSTGRES_HOST=host.docker.internal
POSTGRES_PORT=5432
POSTGRES_DB=mypostgresdb
POSTGRES_USER=mypostgresuser
POSTGRES_PASSWORD=mypostgrespassword
```
4. You should be able to run the server with the following command
``` 
docker compose-up --build
```
