FROM golang:1.17.8

WORKDIR /go/app

ADD . .

RUN go get .

ENV PORT=8080
ENV GIN_MODE="release"

CMD ["go", "run", "main.go"]