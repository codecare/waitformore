# waitformore-frontend

Bootstraped with

    vue create waitformore-frontend

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run buildProd
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

### avoid local npm

start shell in docker using
    
    docker-compose run --rm --service-ports waitformore-npm-dev
    cd workdir
    npm install
    npm run serve
    
### build in docker

    # docker build -t codecare.azurecr.io/waitformore-frontend:0.1.0 .
    docker build -t codecare.azurecr.io/waitformore-frontend:`grep '"version"' package.json | sed -e 's/.*\s"//' -e 's/".*//'` .
    docker push codecare.azurecr.io/waitformore-frontend:`grep '"version"' package.json | sed -e 's/.*\s"//' -e 's/".*//'` 
