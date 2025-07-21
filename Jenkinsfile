pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'banda2133/go-app'
    }

    stages {
       stage('Checkout') {
    steps {
        git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git'
    }
}

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o main .'
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
