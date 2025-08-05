pipeline {
    agent any

    environment {
        HARBOR_URL = 'harbor.example.com'    // Replace with actual Harbor domain
        PROJECT = 'ci-cd'
        IMAGE_NAME = 'go-app'
        TAG = "${env.BUILD_NUMBER}"
    }

    tools {
        go 'go1.24'  // Must match Jenkins Global Tool Configuration name
    }

    stages {

        stage('clone') {
            steps {
                git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git'
            }
        }

        stage('Go Linting') {
            steps {
                sh '''
                    if ! command -v golangci-lint &> /dev/null; then
                      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.58.2
                      export PATH=$(go env GOPATH)/bin:$PATH
                    fi
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
                sh "docker build -t ${HARBOR_URL}/${PROJECT}/${IMAGE_NAME}:${TAG} ."
            }
        }

        stage('Push to Harbor') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'harbor-creds', usernameVariable: 'HARBOR_USER', passwordVariable: 'HARBOR_PASS')]) {
                    sh """
                        echo "${HARBOR_PASS}" | docker login ${HARBOR_URL} -u "${HARBOR_USER}" --password-stdin
                        docker push ${HARBOR_URL}/${PROJECT}/${IMAGE_NAME}:${TAG}
                        docker logout ${HARBOR_URL}
                    """
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
