pipeline {
    agent any

    environment {
        HARBOR_URL = '10.1.0.11:30002'     // 🔁 Replace with your actual Harbor hostname
        PROJECT = 'ci-cd'
        IMAGE_NAME = 'go-app'
        TAG = "${env.BUILD_NUMBER}"
    }

    tools {
        go 'go'  // 🔁 This should match the name in Jenkins > Global Tool Config
    }

    stages {

        stage('Clone Repository') {
            steps {
                git credentialsId: 'github-creds', url: 'https://github.com/ChennuriAkhilvarma/go-app.git', branch: 'master'
            }
        }

        stage('Go Linting') {
    steps {
        sh '''
            export GOPATH=$(go env GOPATH)
            export PATH=$GOPATH/bin:$PATH

            # Install linter if not present
            if ! command -v golangci-lint &> /dev/null; then
                curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.58.2
            fi

            # Run linter with verbose output
            golangci-lint run --verbose --timeout=5m
        '''
    }
}


        stage('Go Build') {
            steps {
                sh 'go mod tidy'
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
