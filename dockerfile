
FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth_service

EXPOSE 8080
CMD ["/auth_service"]