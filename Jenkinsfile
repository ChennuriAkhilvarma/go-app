pipeline {
  agent any

  environment {
    DOCKER_IMAGE = 'banda2133/go-app:latest'
  }

  stages {
    stage('Clone') {
      steps {
        git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git'
      }
    }

    stage('Build Docker Image') {
      steps {
        sh "docker build -t $DOCKER_IMAGE ."
      }
    }

    stage('Push to Docker Hub') {
      steps {
        withCredentials([usernamePassword(
          credentialsId: 'dockerhub-creds',
          usernameVariable: 'DOCKER_USER',
          passwordVariable: 'DOCKER_PWD'
        )]) {
          sh '''
            echo "$DOCKER_PWD" | docker login -u "$DOCKER_USER" --password-stdin
            docker push $DOCKER_IMAGE
          '''
        }
      }
    }
  }

  post {
    failure {
      echo 'Build failed.'
    }
    success {
      echo 'Build succeeded and Docker image pushed.'
    }
  }
}
