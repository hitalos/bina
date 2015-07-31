FROM alpine:latest
MAINTAINER Hítalo Silva <hitalos@jfal.jus.br>

RUN apk update && apk upgrade
RUN apk add php-dom php-json php-ldap php-xml

WORKDIR /var/www
# RUN mkdir classes
# COPY index.php grandstream.php json.php /var/www/
# COPY classes/ldapJFAL.class.php /var/www/classes/

CMD php -S $HOSTNAME:80 index.php
EXPOSE 80
