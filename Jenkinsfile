pipeline {
    agent any

    environment {
        HARBOR_URL = 'harbor.example.com'    // Replace with your actual Harbor hostname (no https)
        PROJECT = 'ci-cd'
        IMAGE_NAME = 'go-app'
        TAG = "${env.BUILD_NUMBER}"
    }

    tools {
        go 'go1.24'   // Matches Jenkins "Go installations" name
    }

    stages {

        stage('Checkout') {
            steps {
                git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git', branch: 'master'
            }
        }

        stage('Go Linting') {
            steps {
                sh '''
                    which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.2
                    golangci-lint run
                '''
            }
        }

        stage('Go Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    sh "docker build -t ${HARBOR_URL}/${PROJECT}/${IMAGE_NAME}:${TAG} ."
                }
            }
        }

        stage('Push to Harbor') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'harbor-creds', usernameVariable: 'HARBOR_USER', passwordVariable: 'HARBOR_PASS')]) {
                    script {
                        sh """
                            echo "${HARBOR_PASS}" | docker login ${HARBOR_URL} -u "${HARBOR_USER}" --password-stdin
                            docker push ${HARBOR_URL}/${PROJECT}/${IMAGE_NAME}:${TAG}
                            docker logout ${HARBOR_URL}
                        """
                    }
                }
            }
        }
    }

    post {
        success {
            echo "✅ Build and push successful: ${HARBOR_URL}/${PROJECT}/${IMAGE_NAME}:${TAG}"
        }
        failure {
            echo "❌ Build failed"
        }
    }
}
