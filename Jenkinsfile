pipeline {
    agent any
    stages {
        stage('Docker Compose up') {
            steps {
                sh 'sudo su docker-compose up -d' 
            }
        }
        stage('Docker Compose logs') {
            steps {
                sh 'sudo su docker-compose logs -f'
            }
        }
        stage('Simple validation'){
            steps {
                sh 'curl -s localhost:8090/api | grep "Jenkins"'
            }
        }
    }
}