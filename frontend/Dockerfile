FROM node:13.10.1-stretch as builder

WORKDIR /client
COPY ./package*.json ./
RUN npm install
COPY . .
COPY .env.production .
RUN npm run buildProd

FROM nginx:1.17.9-alpine
RUN rm -rf /usr/share/nginx/html/*
RUN rm -rf /etc/nginx/conf.d/*
COPY --from=builder /client/dist /usr/share/nginx/html
COPY ./docker/nginx.conf /etc/nginx/conf.d/default.conf