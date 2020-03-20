# Build

    go build cmd/main.go
    
# run

    ./main

# test
    
    curl -v localhost:8081/api/queue -d '{"QueueTitle": "title"}'
    
    curl -v localhost:8081/api/user/NPy8
    curl -v localhost:8081/api/user/NPy8/pull
    
    curl -v localhost:8081/api/queue/3UNd1rmHJUQ99X1v
    curl -v localhost:8081/api/queue/3UNd1rmHJUQ99X1v/call
    
# docker build 

docker build includes go compile!

    docker build -t  wait4more:latest .
    docker build -t 928948368762.dkr.ecr.eu-central-1.amazonaws.com/waitformore-backend:0.1.0 .
    
# azure

    docker login codecare.azurecr.io     
    docker build -t codecare.azurecr.io/waitformore-backend:`head -1 version.txt` .
    docker push codecare.azurecr.io/waitformore-backend:`head -1 version.txt` 

    

    
    