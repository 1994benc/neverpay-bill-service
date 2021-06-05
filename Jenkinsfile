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
        stage("set up environment variables") {
            steps {
                sh 'export POSTGRES_HOST = $POSTGRES_HOST'
                sh 'export POSTGRES_PORT = $POSTGRES_PORT'
                sh 'export POSTGRES_DB = $POSTGRES_DB'
                sh 'export POSTGRES_USER = $POSTGRES_USER'
                sh 'export POSTGRES_PASSWORD = $POSTGRES_PASSWORD'
                sh 'export SSL_MODE = $SSL_MODE'
            }
        }
        stage("build") {
            steps {
                echo "Building the docker containers..."
                sh '/usr/local/bin/docker-compose build'
            }
        }
        stage("up") {
            steps {
                echo "Spinning up the docker containers..."
                sh '/usr/local/bin/docker-compose up'
            }
        }
    }
}