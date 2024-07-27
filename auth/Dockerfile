FROM golang:1.22.4-alpine
WORKDIR /auth
COPY . .
RUN go mod download
RUN go build -o auth
EXPOSE 3000
CMD ["./auth"]
