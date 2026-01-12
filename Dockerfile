## Staging Build
FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN apk update 
RUN apk add tzdata
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o fraud-detection-engine cmd/service/main.go

## Build after staging
FROM alpine:latest

WORKDIR /app
ENV TZ=Asia/Jakarta
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/fraud-detection-engine /app/fraud-detection-engine

COPY .env.sample .env

ENTRYPOINT ["/app/fraud-detection-engine"]