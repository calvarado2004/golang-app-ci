FROM docker.io/bitnami/golang:1.16.15 AS builder
ADD ./src /app
WORKDIR /app
RUN go mod init github.com/calvarado2004/app-golang && go mod tidy && go get github.com/cosmtrek/air && go get github.com/gofiber/template/html
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo server.go

FROM busybox:stable
ENV APP_HOME /app
RUN adduser 1001 -D -h $APP_HOME && mkdir -p $APP_HOME && chown 1001:1001 $APP_HOME
USER 1001
WORKDIR $APP_HOME
COPY ./src/views views/
COPY ./src/static static/
COPY --chown=0:0 --from=builder /app/server ./
EXPOSE 8080
CMD ["./server"]
