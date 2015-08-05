FROM alpine:latest
MAINTAINER Hítalo Silva <hitalos@jfal.jus.br>

RUN apk update && apk upgrade
RUN apk add php-dom php-json php-ldap php-xml

VOLUME /var/www
WORKDIR /var/www
CMD php -S 0.0.0.0:80 index.php
EXPOSE 80
