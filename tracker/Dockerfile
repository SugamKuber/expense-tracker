FROM golang:1.22.4-alpine
WORKDIR /tracker
COPY . .
RUN go mod tidy
RUN go build -o tracker ./cmd/main.go
EXPOSE 3001
CMD ["./tracker"]