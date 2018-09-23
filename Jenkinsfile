pipeline {
    agent any 
    stages {
        stage('Test') {
            steps {
                echo 'Hello world!' 
            }
        }
        stage('Test Master') {
            when {
                branch 'master'
            }
            steps {
                echo 'Master branch!' 
            }
        }
    }
}