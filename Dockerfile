FROM golang:1.19.1 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app cmd/main.go

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 8080
CMD ./app