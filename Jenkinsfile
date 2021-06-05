pipeline {
    agent any
    environment {
            POSTGRES_HOST = credentials('POSTGRES_HOST')
            POSTGRES_PORT = credentials('POSTGRES_PORT')
            POSTGRES_DB = credentials('POSTGRES_DB')
            POSTGRES_USER = credentials('POSTGRES_USER')
            POSTGRES_PASSWORD = credentials('POSTGRES_PASSWORD')
            SSL_MODE = credentials('SSL_MODE')
    }
    stages {
        
        stage("build") {
            steps {
                echo "Building the docker containers..."
                sh '/usr/local/bin/docker-compose build'
            }
        }
        stage("up") {
            steps {
                echo "Spinning up the docker containers..."
                sh '/usr/local/bin/docker-compose up -d --verbose'
            }
        }
    }
}