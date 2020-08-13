pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '/var/jenkins_home/workspace/script/build.sh'
      }
    }

    stage('Test') {
      steps {
        sh '/var/jenkins_home/workspace/script/test.sh'
      }
    }

    stage('Deploy') {
      steps {
        sh '/var/jenkins_home/workspace/script/deploy.sh'
      }
    }

  }
}