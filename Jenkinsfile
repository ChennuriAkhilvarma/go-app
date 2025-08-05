pipeline {
    agent any

    environment {
        IMAGE_NAME = "go-app"
        HARBOR_URL = "harbor.example.com"      // 🔁 Replace with your Harbor domain
        PROJECT = "library"                    // 🔁 Replace with your Harbor project name
        TAG = "${BUILD_NUMBER}"
    }

    stages {

        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }

        stage('Lint') {
            steps {
                echo "Running golangci-lint..."
                sh '''
                    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
                    $(go env GOPATH)/bin/golangci-lint run
                '''
            }
        }

        stage('Unit Tests') {
            steps {
                echo "Running go test..."
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
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
        failure {
            echo '❌ Build failed!'
        }
        success {
            echo '✅ Build and push to Harbor successful!'
        }
    }
}
