FROM node:alpine
LABEL maintainer="Hítalo Silva <hitalos@gmail.com>"

RUN apk update && apk upgrade

ADD . /app
WORKDIR /app

# Dependencies to build libxmljs
ENV DEV_LIBS 'g++ gcc libxml2-dev make python'

RUN apk add $DEV_LIBS && \
  yarn && \
  yarn run build && \
  rm -rf node_modules && \
  yarn --prod && \
  yarn cache clean && \
  apk del $DEV_LIBS

CMD npm start
