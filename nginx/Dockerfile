FROM ubuntu:24.04

WORKDIR /nginx

COPY . .

RUN apt-get update && \
    apt-get install -y nginx && \
    cp nginx.conf /etc/nginx/nginx.conf

CMD ["nginx","-g","daemon off;"]

