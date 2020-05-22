# waitformore

wait for more is a queueing service which allows anyone to provide a waiting queue

see a running version on:

    https://wartewarte.de
    ssh waitformore@wartewarte.de

# deployment

### containers
we use 3 containers for the workload and a reverse proxy

- backend 
- frontend
- database

### backend

build backend using docker
see backend/README.md
backend is written in golang 

### frontend

build frontend using docker
see frontend/README.md
frontend is written in vuejs with vuetify

### continuous deployment with Jenkinsfile

jenkins needs some credentials (see Jenkinsfile)




