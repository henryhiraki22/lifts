pipeline {
    agent any
    stages {
        stage('Docker stop') {
            steps {
                sh 'sudo docker stop $(docker ps -a -q)'
            }
        }
        stage('Docker rm') {
            steps {
                sh 'sudo docker rm $(docker ps -a -q)'
            }
        }
        stage('Docker rmi') {
            steps {
                sh 'sudo docker rmi $(docker images -q)'
            }
        }
        stage('Docker Compose up') {
            steps {
                sh 'sudo docker-compose up -d' 
            }
        }
        stage('Docker Compose logs') {
            steps {
                sh 'sudo docker-compose logs -f'
            }
        }
        stage('Simple validation'){
            steps {
                sh 'curl -s localhost:8090/api | grep "Jenkins"'
            }
        }
    }
}