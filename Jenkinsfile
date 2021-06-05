pipeline {
    agent any
    environment {
            ENV_FILE = credentials('env')
    }
    stages {
        stage("build") {
            steps {
                echo "Building the docker containers..."
                sh '/usr/local/bin/docker-compose --env-file $ENV_FILE build'
            }
        }
        stage("up") {
            steps {
                echo "Spinning up the docker containers..."
                sh '/usr/local/bin/docker-compose --env-file $ENV_FILE up'
            }
        }
    }
}