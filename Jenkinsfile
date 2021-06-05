pipeline {
    agent any
    environment {
            ENV_FILE = credentials('env')
    }
    stages {
        
        stage("build") {
            steps {
                echo "Building the docker containers..."
                sh "/usr/local/bin/docker-compose build --env-file $ENV_FILE" 
            }
        }
        stage("up") {
            steps {
                echo "Spinning up the docker containers..."
                sh "/usr/local/bin/docker-compose up -d --env-file $ENV_FILE" 
            }
        }
    }
}