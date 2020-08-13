pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '/usr/local/go/bin/go build -v -o orderlunch.bin'
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