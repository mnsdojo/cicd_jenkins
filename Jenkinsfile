pipeline {
    agent any
     triggers {
        pollSCM('H/5 * * * *')  // 👈 add this
    }
    environment {
        GOPATH = "${WORKSPACE}/go"
        PATH = "${env.PATH}:${WORKSPACE}/go/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }
        stage('Test'){
            steps {
                sh 'go test'
            }
        }
        stage('Build') {
            steps {
                sh 'go version'
                sh 'go build -o app .'
            }
        }

        stage('Package') {
            steps {
                sh 'tar -czf app.tar.gz app'
            }
        }

    }

    post {
        sucess {
            echo 'Build Passed!'
        }
        failure {
            echo 'Build failed'
        }
        always {
            echo "Pipeline finished Workspace : ${WORKSPACE}"
        }
    }
}