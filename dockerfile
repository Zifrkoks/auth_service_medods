FROM golang:1.24.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/swaggo/gin-swagger
RUN go install github.com/swaggo/files
RUN swag init -g main.go -o ./docs
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./main.go

EXPOSE 8080
CMD ["/main"]