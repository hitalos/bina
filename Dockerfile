FROM node:13-alpine as frontend-builder
WORKDIR /app
ADD package.json webpack.config.js ./
ADD src/ src/
RUN npm i && NODE_ENV=production npm run build

FROM golang:alpine as backend-builder
WORKDIR /app
ADD go.mod go.sum ./
COPY --from=frontend-builder /app/public/ .
RUN go get github.com/GeertJohan/go.rice && \
	go get github.com/GeertJohan/go.rice/rice
ADD . .
RUN rice -i ./cmd embed-go
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o bina ./cmd

FROM scratch
LABEL maintainer="Hítalo Silva <hitalos@gmail.com>"
WORKDIR /app
COPY --from=backend-builder /app/bina .

ENTRYPOINT ["/app/bina"]
