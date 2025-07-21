pipeline {
    agent {
        docker {
            image 'golang:1.24.2' // or any version that has Go
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    environment {
        DOCKER_IMAGE = 'banda2133/go-app'
    }

    stages {
        stage('Check Tools') {
            steps {
                sh 'go version'
                sh 'docker --version'
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

        stage('Docker Build & Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh 'docker build -t $DOCKER_IMAGE:latest .'
                    sh 'echo "$PASSWORD" | docker login -u "$USERNAME" --password-stdin'
                    sh 'docker push $DOCKER_IMAGE:latest'
                }
            }
        }
    }
}
