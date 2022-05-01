FROM docker.io/bitnami/golang:1.16.15 AS builder
ADD ./src /app
WORKDIR /app
RUN go mod init test/m && go mod tidy && go get github.com/cosmtrek/air && go get github.com/gofiber/template/html
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo server.go

FROM busybox:stable
ENV APP_HOME /app
RUN wget https://github.com/moparisthebest/static-curl/releases/download/v7.83.0/curl-amd64 && mv curl-amd64 /usr/local/bin/curl && chmod 755 /usr/local/bin/curl && adduser 1001 -D -h $APP_HOME && mkdir -p $APP_HOME && chown 1001:1001 $APP_HOME
USER 1001
WORKDIR $APP_HOME
COPY ./src/views views/
COPY --chown=0:0 --from=builder /app/server ./
EXPOSE 8080
CMD ["./server"]