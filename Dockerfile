FROM node:alpine
MAINTAINER Hítalo Silva <hitalos@gmail.com>

RUN apk update && apk upgrade

ADD . /app
WORKDIR /app

RUN yarn && NODE_ENV=prod yarn run build && rm -rf node_modules && yarn --prod && yarn cache clean

CMD npm start
