pipeline {
    agent {
        docker {
            image 'golang:1.24.2'
            args '-v /var/run/docker.sock:/var/run/docker.sock' // To enable Docker-in-Docker
        }
    }

    environment {
        DOCKER_IMAGE = 'banda2133/go-app' // Replace if your Docker Hub username or repo differs
    }

    stages {
        stage('Docker Check') {
            steps {
                sh 'docker --version'
                sh 'docker ps'
            }
        }

        stage('Checkout') {
            steps {
                git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o app'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE:latest .'
            }
        }

        stage('Docker Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh 'echo "$PASSWORD" | docker login -u "$USERNAME" --password-stdin'
                    sh 'docker push $DOCKER_IMAGE:latest'
                }
            }
        }
    }
}
