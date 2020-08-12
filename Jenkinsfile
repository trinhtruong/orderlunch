pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'sudo apt-get update; wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz;sudo tar -xvf go1.13.3.linux-amd64.tar.gz;sudo mv go /usr/local; '
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