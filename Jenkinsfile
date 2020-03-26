pipeline {
    agent { label 'TRAX-JDK10' }

    stages {

        stage('Build Backend') {

            steps {
                ansiColor('xterm') {
                    dir('backend') {
                        sh 'docker build -t codecare.azurecr.io/waitformore-backend:`head -1 version.txt` .'
                    }
                }
            }
        }
        stage('Docker Login') {
            when {
                when { anyOf { branch 'deploy-prod-azure';branch 'deploy-test-azure' } }
            }
            steps {
                ansiColor('xterm') {
                    dir('backend') {
                        withCredentials([usernamePassword(credentialsId: 'codecare.azurecr.io.login', passwordVariable: 'DOCKERPASS', usernameVariable: '')]) {
                            sh '''
                                docker login codecare.azurecr.io -u codecare -p "$DOCKERPASS"
                            '''
                        }
                    }
                }
            }
        }
        stage('Push Backend') {
            when {
                when { anyOf { branch 'deploy-prod-azure';branch 'deploy-test-azure' } }
            }
            steps {
                ansiColor('xterm') {
                    dir('backend') {
                        sh '''
                            docker push codecare.azurecr.io/waitformore-backend:`head -1 version.txt`
                        '''
                    }
                }
            }
        }
        stage('Build Frontend') {

            steps {
                ansiColor('xterm') {
                    dir('frontend') {
                        sh '''
                            FRONTEND_VERSION=`grep '"version"' package.json | sed -e 's/.*\\s"//' -e 's/".*//'`
                            docker build -t codecare.azurecr.io/waitformore-frontend:${FRONTEND_VERSION} .
                        '''
                    }
                }
            }
        }

        stage('Push Frontend') {
            when {
                when { anyOf { branch 'deploy-prod-azure';branch 'deploy-test-azure' } }
            }
            steps {
                ansiColor('xterm') {
                    dir('frontend') {
                        sh '''
                            FRONTEND_VERSION=`grep '"version"' package.json | sed -e 's/.*\\s"//' -e 's/".*//'`
                            docker push codecare.azurecr.io/waitformore-frontend:${FRONTEND_VERSION}
                        '''
                    }
                }
            }
        }

        stage('Ansible Deploy Prod') {
            when {
                branch 'deploy-prod-azure'
            }
            steps {
                ansiColor('xterm') {
                    dir('ansible/prod') {
                        withCredentials(
                            [sshUserPrivateKey(credentialsId: 'jenkins-for-waitforemore-ssh', keyFileVariable: 'ANSIBLE_PRIVATE_KEY', passphraseVariable: '', usernameVariable: 'ANSIBLE_USERNAME')],
                            [usernamePassword(credentialsId: 'codecare.azure.waitformore.prod.db_pass', passwordVariable: 'DB_PASSWORD', usernameVariable: '')]
                            ) {
                            sh '''
                                FRONTEND_VERSION=`grep '"version"' ../../frontend/package.json | sed -e 's/.*\\s"//' -e 's/".*//'`
                                BACKEND_VERSION=`head -1 ../../backend/version.txt`
                                echo "versions found: $BACKEND_VERSION + $FRONTEND_VERSION"
                                # known host
                                ssh-keyscan -H wartewarte.de >> /home/jenkins/.ssh/known_hosts
                                ansible-playbook -i waitformoredev01 deploy.yml --private-key ${ANSIBLE_PRIVATE_KEY} --extra-vars "database_pwd=$DB_PASSWORD backend_version=$BACKEND_VERSION frontend_version=$FRONTEND_VERSION"
                                # remove known host
                                ssh-keygen -R wartewarte.de

                            '''
                        }
                    }
                }
            }
        }

        stage('Ansible Deploy Test') {
            when {
                branch 'deploy-test-azure'
            }
            steps {
                ansiColor('xterm') {
                    dir('ansible/test') {
                        withCredentials(
                            [sshUserPrivateKey(credentialsId: 'jenkins-for-waitforemore-ssh', keyFileVariable: 'ANSIBLE_PRIVATE_KEY', passphraseVariable: '', usernameVariable: 'ANSIBLE_USERNAME')],
                            [usernamePassword(credentialsId: 'codecare.azure.waitformore.test.db_pass', passwordVariable: 'DB_PASSWORD', usernameVariable: '')]
                            ) {
                            sh '''
                                FRONTEND_VERSION=`grep '"version"' ../../frontend/package.json | sed -e 's/.*\\s"//' -e 's/".*//'`
                                BACKEND_VERSION=`head -1 ../../backend/version.txt`
                                echo "versions found: $BACKEND_VERSION + $FRONTEND_VERSION"
                                # known host
                                ssh-keyscan -H wartewarte.de >> /home/jenkins/.ssh/known_hosts
                                ansible-playbook -i waitformoredev01 deploy.yml --private-key ${ANSIBLE_PRIVATE_KEY} --extra-vars "database_pwd=$DB_PASSWORD backend_version=$BACKEND_VERSION frontend_version=$FRONTEND_VERSION"
                                # remove known host
                                ssh-keygen -R wartewarte.de

                            '''
                        }
                    }
                }
            }
        }
    }
}
