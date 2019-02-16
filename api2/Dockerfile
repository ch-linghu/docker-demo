FROM golang:alpine

WORKDIR /go/src/app
COPY main.go .

RUN go install -v ./...

CMD ["app"]
