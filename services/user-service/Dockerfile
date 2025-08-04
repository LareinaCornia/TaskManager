FROM golang:1.22-apline

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o user-service main.go

EXPOSE 8081
CMD ["./user-service"]