pipeline {
  agent any

 environment {
        DOCKER_IMAGE = 'banda2133/go-app'
    }

  stages {
    stage('Clone') {
      steps {
        git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git'
      }
    }

    stage('Build Docker Image') {
      steps {
        sh 'docker build -t banda2133/go-app:latest .'
      }
    }

    stage('Push to Docker Hub') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', passwordVariable: 'DOCKER_PWD', usernameVariable: 'DOCKER_USER')]) {
          sh '''
            echo "$DOCKER_PWD" | docker login -u "$DOCKER_USER" --password-stdin
            docker push banda2133/go-app:latest
          '''
        }
      }
    }
  }
}












    

    

    