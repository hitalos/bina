FROM alpine:latest
MAINTAINER Hítalo Silva <hitalos@jfal.jus.br>

RUN apk update && apk upgrade
RUN apk add php-dom php-json php-ldap php-phar php-openssl php-xml

VOLUME /var/www
WORKDIR /var/www
COPY . /var/www

CMD php -S 0.0.0.0:80 -t public public/index.php

ENV VIRTUAL_HOST=bina.jfal.jus.br
EXPOSE 80
