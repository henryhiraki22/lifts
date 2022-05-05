pipeline {
    agent any
    stage('Docker stop') {
        steps {
            sh 'docker stop $(docker ps -a -q)'
        }
    }
    stage('Docker rm') {
        steps {
            sh 'docker rm $(docker ps -a -q)'
        }
    }
    stage('Docker rmi') {
        steps {
            sh 'docker rmi $(docker images -q)'
        }
    }
    stage('Docker Compose up') {
        steps {
            sh 'docker-compose up -d' 
        }
    }
    stage('Docker Compose logs') {
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