FROM docker.io/bitnami/golang:1.16.15
ADD ./src /app
WORKDIR /app
RUN go mod init test/m && go mod tidy && go get github.com/cosmtrek/air && go get github.com/gofiber/template/html
EXPOSE 8080
CMD ["go", "run", "github.com/cosmtrek/air"]
