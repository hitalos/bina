FROM docker.io/library/node:16-alpine as frontend-builder
WORKDIR /app
COPY package.json package-lock.json build.js ./
COPY src/ src/
RUN npm i && NODE_ENV=production npm run build

FROM docker.io/library/golang:alpine as backend-builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
COPY --from=frontend-builder /app/cmd/public/scripts cmd/public/scripts
COPY --from=frontend-builder /app/cmd/public/styles cmd/public/styles
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o bina ./cmd

FROM scratch
LABEL maintainer="Hítalo Silva <hitalos@gmail.com>"
WORKDIR /app
COPY --from=backend-builder /app/bina .

ENTRYPOINT ["/app/bina"]
