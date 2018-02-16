FROM ubuntu:latest

COPY image.png /www
COPY nginx.conf /etc/nginx/nginx.conf

RUN apt-get update && apt-get install -y \
        libpcre3 libpcre3-dev \
        zlib1g-dev \
        libssl-dev

RUN apt-get install -y git

RUN git clone https://github.com/nginx/nginx \
    && cd nginx \
    && ./auto/configure

RUN ./auto/configure \
    --with-http_v2_module \
    --with-http_ssl_module

RUN make -j4

EXPOSE 443
