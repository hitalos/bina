FROM docker.io/library/node:22-alpine as frontend-builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci && mkdir -p cmd/public/assets/scripts cmd/public/assets/styles
COPY cmd/public/assets/ cmd/public/assets
RUN NODE_ENV=production npm run build::js && npm run build::css

FROM docker.io/library/golang:1.25.3-alpine as backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && mkdir -p cmd/public/assets/scripts cmd/public/assets/styles
COPY . .
COPY --from=frontend-builder /app/cmd/public/assets/scripts cmd/public/assets/scripts
COPY --from=frontend-builder /app/cmd/public/assets/styles cmd/public/assets/styles
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o bina ./cmd

FROM scratch
LABEL maintainer="Hítalo Silva <hitalos@gmail.com>"
WORKDIR /app
COPY --from=backend-builder /app/bina .

ENTRYPOINT ["/app/bina"]
