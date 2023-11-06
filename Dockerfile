FROM golang:latest AS builder
WORKDIR /app
COPY  . .
RUN chmod 777 *
RUN  go build -o webhook-siscope
RUN ls -la

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/webhook-siscope /app
LABEL authors="edwin"
EXPOSE 8090

ENTRYPOINT ["./webhook-siscope"]
