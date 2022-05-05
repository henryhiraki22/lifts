pipeline {
    agent any
    stages {
        stage('Docker Compose up') {
            steps {
                sh 'docker-compose up -d' 
            }
        }
        stage('Docker Compose log') {
            steps {
                sh 'docker-compose logs -f'
            }
        }
        stage('Simple validation'){
            steps {
                sh 'curl -s localhost:8090/api | grep "Jenkins"'
            }
        }
    }
}