pipeline {
  environment {
    registry = "airren/echo-bio-cloud"
    registryCredential = ‘dockerhub’
  }
  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git 'https://github.com/Airren/echo-bio-cloud.git'
      }
    }
    stage('Building image') {
      steps{
        script {
          docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }
//       stage('Deploy Image') {
//           steps{
//             script {
//               docker.withRegistry( '', registryCredential ) {
//                 dockerImage.push()
//               }
//             }
//           }
//         }
//
        stage('Docker Push') {
            	agent any
              steps {
              	withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'Forever8023_', usernameVariable: 'airren')]) {
                	sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                  sh 'docker push shanem/spring-petclinic:latest'
                }
              }
            }
        stage('Remove Unused docker image') {
          steps{
            sh "docker rmi $registry:$BUILD_NUMBER"
          }
        }
  }
}