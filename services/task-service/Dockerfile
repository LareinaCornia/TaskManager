FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o task-service main.go

EXPOSE 8082

CMD ["./task-service"]