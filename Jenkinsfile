pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo "Building the docker containers..."
                sh "sudo docker-compose build" 
            }
        }
        stage("up") {
            steps {
                echo "Spinning up the docker containers..."
                sh "sudo docker-compose up" 
            }
        }
    }
}