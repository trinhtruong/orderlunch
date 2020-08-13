pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'export GOROOT=/usr/local/go'
        sh 'export PATH=$GOROOT/bin:$PATH; go get -u github.com/gorilla/mux'
        sh 'go get -u github.com/lib/pq'
        sh 'go build -v -o orderlunch.bin'
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