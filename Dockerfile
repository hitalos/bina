FROM node:alpine
MAINTAINER Hítalo Silva <hitalos@gmail.com>

RUN apk update && apk upgrade

ADD . /app
WORKDIR /app

# Dependencies to build libxmljs
ENV DEV_LIBS 'g++ gcc libxml2-dev make python'

RUN apk add $DEV_LIBS && \
  yarn && \
  NODE_ENV=prod yarn run build && \
  rm -rf node_modules && \
  yarn --prod && \
  apk del $DEV_LIBS && \
  yarn clean && \
  yarn cache clean

CMD npm start
