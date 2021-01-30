FROM golang:1.13.8
WORKDIR /go/src/app
COPY . .
RUN go build -o main -i ./cmd
CMD ["./main","-file=dev_local.json"]