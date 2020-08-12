pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'git pull origin master'
      }
    }

    stage('Test') {
      steps {
        sh 'go test -v'
      }
    }

    stage('Deploy') {
      steps {
        sh 'go run main.go'
      }
    }

  }
}