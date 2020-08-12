pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'apt-get update;wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz;tar -xvf go1.13.3.linux-amd64.tar.gz;mv go /usr/local; '
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