pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''
export GOROOT=/usr/local/go
bash-4.4;export PATH=$GOROOT/bin:$PATH;'''
        sh 'go get -u github.com/gorilla/mux;go get -u github.com/lib/pq;'
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