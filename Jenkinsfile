pipeline {
    agent any
    options {
        buildDiscarder(logRotator(
            numToKeepStr: '10',
            daysToKeepStr: '30',
            artifactNumToKeepStr: '5',
            artifactDaysToKeepStr: '7'
        ))
        timestamps()
        timeout(time: 10, unit: 'MINUTES')
    }
    triggers {
        pollSCM('H/5 * * * *')
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
        stage('Test') {
            steps {
                sh 'go test ./...'
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
                archiveArtifacts artifacts: 'app.tar.gz',
                                 onlyIfSuccessful: true
            }
        }
    }
    post {
        success {
            echo 'Build Passed!'
        }
        failure {
            echo 'Build failed'
        }
        always {
            cleanWs()
            echo "Pipeline finished Workspace: ${WORKSPACE}"
        }
    }
}