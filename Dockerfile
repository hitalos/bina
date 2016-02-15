FROM alpine:latest
MAINTAINER Hítalo Silva <hitalos@jfal.jus.br>

RUN apk update && apk upgrade
RUN apk add git nodejs php-dom php-json php-ldap php-phar php-openssl php-xml

RUN npm install -g bower
RUN php -r "readfile('https://getcomposer.org/installer');" | php && \
	mv composer.phar /usr/bin/composer && \
	chmod +x /usr/bin/composer
RUN
VOLUME /var/www
WORKDIR /var/www
COPY . /var/www

CMD bower install --allow-root && \
    composer -n --no-progress install && \
    php -S 0.0.0.0:80 -t public public/index.php

ENV VIRTUAL_HOST=bina.jfal.jus.br
EXPOSE 80
