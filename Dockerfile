FROM node:10-alpine as builder
LABEL maintainer="Hítalo Silva <hitalos@gmail.com>"

ADD package.json /app/
WORKDIR /app

# Dependencies to build libxmljs
ENV DEV_LIBS 'g++ gcc libxml2-dev make python'

RUN apk -U add $DEV_LIBS
RUN yarn
ENV NODE_ENV production
COPY package.json gulpfile.js ./
COPY src/ ./src/
COPY public/ ./public/
COPY bin/ ./bin/

RUN npm run build
RUN rm -rf node_modules && yarn

# production image
FROM node:10-alpine

WORKDIR /app
COPY --from=builder /app .

EXPOSE 3000
CMD npm start
