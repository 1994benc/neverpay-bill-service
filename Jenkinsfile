pipeline {
    agent any
    environment {
            JPOSTGRES_HOST = credentials('POSTGRES_HOST')
            JPOSTGRES_PORT = credentials('POSTGRES_PORT')
            JPOSTGRES_DB = credentials('POSTGRES_DB')
            JPOSTGRES_USER = credentials('POSTGRES_USER')
            JPOSTGRES_PASSWORD = credentials('POSTGRES_PASSWORD')
            JSSL_MODE = credentials('SSL_MODE')
    }
    stages {
        stage("set up environment variables") {
            steps {
                sh 'export POSTGRES_HOST = $JPOSTGRES_HOST'
                sh 'export POSTGRES_PORT = $JPOSTGRES_PORT'
                sh 'export POSTGRES_DB = $JPOSTGRES_DB'
                sh 'export POSTGRES_USER = $JPOSTGRES_USER'
                sh 'export POSTGRES_PASSWORD = $JPOSTGRES_PASSWORD'
                sh 'export SSL_MODE = $JSSL_MODE'
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