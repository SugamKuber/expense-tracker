FROM golang:1.22.4-alpine
WORKDIR /tracker
COPY . .
RUN go mod download
RUN go build -o tracker
EXPOSE 3001
CMD ["./tracker"]
