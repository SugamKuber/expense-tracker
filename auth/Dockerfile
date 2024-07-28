FROM golang:1.22.4-alpine
WORKDIR /auth
COPY . .
RUN go mod tidy
RUN go build -o auth ./cmd/main.go
EXPOSE 3000
CMD ["./auth"]
