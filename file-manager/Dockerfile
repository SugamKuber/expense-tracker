FROM golang:1.22.4-alpine
WORKDIR /file-manager
COPY . .
RUN go mod tidy
RUN go build -o file-manager ./cmd/main.go
EXPOSE 3002
CMD ["./file-manager"]