FROM ubuntu:24.04

WORKDIR /nginx

COPY . .

RUN apt-get update && \
    apt-get install -y --no-install-recommends nginx && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/doc/* /usr/share/man/*

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8081 8082

CMD ["nginx","-g","daemon off;"]

