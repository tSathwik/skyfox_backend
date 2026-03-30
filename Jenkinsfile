pipeline {
    agent any 

    environment {
        IMAGE_NAME = 'skyfox_backend'
        CONTAINER_NAME = 'skyfox_backend'
        PORT = '8000'
    }


    stages {

        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Run Tests') {
            steps {
                sh '''
                docker run --rm \
                  -v $(pwd):/app \
                  -w /app \
                  golang:1.25-alpine \
                  go test ./...
                '''
            }
        }

        stage('Build Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

        stage('Remove Old Container') {
            steps {
                sh '''
                docker rm -f $CONTAINER_NAME || true
                '''
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                docker compose down || true
                docker compose up -d --build
                '''
            }
        }
    }
}