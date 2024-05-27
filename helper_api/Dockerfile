FROM golang:1.22-alpine as go-builder
WORKDIR /app
COPY . .
RUN apk add -u -t build-tools curl git && \
    go build -o server *.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*

FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=go-builder /app/server /app/server
COPY ./debian.sh /app
EXPOSE 8080
ENTRYPOINT [ "/app/server" ]