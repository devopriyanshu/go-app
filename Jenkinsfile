pipeline {
    agent none

    environment {
        SONAR_HOST_URL = "http://34.31.6.209:9000"
        IMAGE = "devopriyanshu/go-app:latest"
    }

    stages {

        stage('Checkout') {
            agent any
            steps {
                git branch: 'main', url: 'https://github.com/devopriyanshu/go-app.git'
            }
        }

        stage('Test (Go inside Docker)') {
            agent {
                docker {
                    image 'golang:1.22-alpine'
                    args '-v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            steps {
                sh '''
                export GOCACHE=/tmp/go-cache
                mkdir -p /tmp/go-cache
                go test ./... -coverprofile=coverage.out
                '''
            }
        }

        stage('Build Go Binary (Go inside Docker)') {
            agent {
                docker {
                    image 'golang:1.22-alpine'
                    args '-v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            steps {
                sh '''
                export GOCACHE=/tmp/go-cache
                mkdir -p /tmp/go-cache
                go build -o app
                '''
            }
        }

        stage('Sonar Scan') {
            agent any
            steps {
                withCredentials([string(credentialsId: 'sonar-token', variable: 'SONAR_AUTH_TOKEN')]) {
                    sh """
                    sonar-scanner \
                      -Dsonar.projectKey=go-demo \
                      -Dsonar.host.url=${SONAR_HOST_URL} \
                      -Dsonar.login=$SONAR_AUTH_TOKEN \
                      -Dsonar.go.coverage.reportPaths=coverage.out
                    """
                }
            }
        }

        stage('Build & Push Image') {
            agent any
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'docker-cred',
                    usernameVariable: 'DOCKERHUB_USER',
                    passwordVariable: 'DOCKERHUB_PASS'
                )]) {
                    sh """
                    echo "$DOCKERHUB_PASS" | docker login -u "$DOCKERHUB_USER" --password-stdin
                    docker build -t $IMAGE .
                    docker push $IMAGE
                    """
                }
            }
        }

        stage('Deploy') {
            agent any
            steps {
                sh """
                docker pull $IMAGE
                docker stop go-app || true
                docker rm go-app || true
                docker run -d --name go-app -p 8081:8081 $IMAGE
                """
            }
        }
    }
}
