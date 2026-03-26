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

        stage('Build') {
            steps {
                sh 'go build -o app .'
            }
        }

        stage('Package') {
            steps {
                sh 'tar -czf app.tar.gz app'
            }
        }
    }
}