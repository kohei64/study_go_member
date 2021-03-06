FROM golang:1.18.1-alpine

ENV ROOT=/go/src
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:en
ENV LC_ALL ja_JP.UTF-8

WORKDIR ${ROOT}

COPY . ${ROOT}

RUN go mod download

EXPOSE 8080

RUN go build main.go

CMD ["go", "run", "main.go"]
