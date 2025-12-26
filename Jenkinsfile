pipeline {
    agent any

    environment {
        APP_NAME = "storage-api"
        CONTAINER_NAME = "REKA-Service-Storage"
    }

    stages {

        stage('Checkout') {
            steps {
                echo "🔄 Checkout source code"
                checkout scm
            }
        }

        stage('Stop Old Container') {
            steps {
                echo "🛑 Stop and remove old container if exists"
                sh '''
                    set -x
                    docker ps -a | grep ${CONTAINER_NAME} && docker stop ${CONTAINER_NAME} || true
                    docker ps -a | grep ${CONTAINER_NAME} && docker rm -f ${CONTAINER_NAME} || true
                '''
            }
        }

        stage('Build Image') {
            steps {
                echo "🏗 Build Docker image (no cache)"
                sh '''
                    set -x
                    # Build image tanpa copy .env
                    docker compose -f docker-compose.yml build --no-cache
                '''
            }
        }

        stage('Deploy') {
            steps {
                echo "🚀 Deploy container"
                sh '''
                    set -x
                    # Gunakan env_file host sehingga container dapat environment
                    docker compose -f docker-compose.yml up -d
                '''
            }
        }

        stage('Verify') {
            steps {
                echo "✅ Verify deployment"
                sh '''
                    set -x
                    docker ps | grep ${CONTAINER_NAME} || echo "Container not running!"
                    echo "==== Container Logs ===="
                    docker logs ${CONTAINER_NAME} || true
                '''
            }
        }
    }

    post {
        success {
            echo '✅ DEPLOY SUCCESS'
        }
        failure {
            echo '❌ DEPLOY FAILED'
        }
    }
}
