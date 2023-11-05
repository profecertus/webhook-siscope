FROM golang:latest AS builder
WORKDIR /app
COPY main.go /app
RUN /bin/sh -c go build -o webhook-siscope

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/webhook-siscope /app
LABEL authors="edwin"
EXPOSE 8090

ENTRYPOINT ["webhook-siscope"]