FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o notification-service main.go

EXPOSE 8083

CMD ["./notification-service"]