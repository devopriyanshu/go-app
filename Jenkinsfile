pipeline {
    agent {
        docker {
            image 'golang:1.22-alpine'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    environment {
        SONAR_HOST_URL = "http://34.31.6.209:9000"
        SONAR_TOKEN = credentials('sonar-token')
        IMAGE = "devopriyanshu/go-app:latest"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/devopriyanshu/go-app.git'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -coverprofile=coverage.out'
            }
        }

        stage('Build Go Binary') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Sonar Scan') {
            steps {
                sh """
                sonar-scanner \
                  -Dsonar.projectKey=go-demo \
                  -Dsonar.host.url=$SONAR_HOST_URL \
                  -Dsonar.login=$SONAR_TOKEN \
                  -Dsonar.go.coverage.reportPaths=coverage.out
                """
            }
        }

        stage('Build & Push Image') {
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
