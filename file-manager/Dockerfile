FROM golang:1.22.4-alpine
WORKDIR /file-manager
COPY . .
RUN go mod download
RUN go build -o file-manager
EXPOSE 3002
CMD ["./file-manager"]
