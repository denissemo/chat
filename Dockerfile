FROM golang:1.15-alpine

RUN export GOPATH=/go

RUN go env GOPATH

COPY . /go/src/chat

WORKDIR /go/src/chat

CMD [ "go", "run", "main.go" ]
