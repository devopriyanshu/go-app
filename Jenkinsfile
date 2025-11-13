pipeline {
    agent any

    environment {
        SONAR_HOST_URL = "http://34.31.6.209/9000"
        SONAR_TOKEN = credentials('sonar-token')
        IMAGE = "devopriyanshu/go-app:latest"
    }

    stages {

        stage('Checkout') {
            steps {
                git 'https://github.com/devopriyanshu/go-app.git'
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
                    echo "Logging into Docker Hub..."
                    echo "$DOCKERHUB_PASS" | docker login -u "$DOCKERHUB_USER" --password-stdin

                    echo "Building Docker image..."
                    docker build -t $IMAGE .

                    echo "Pushing Docker image..."
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
