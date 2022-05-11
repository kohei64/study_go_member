# as○○とは
# alpineとは
FROM golang:1.18.1

ENV ROOT=/go/src
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR ${ROOT}

COPY . ${ROOT}

RUN go mod download

EXPOSE 8080

RUN go build main.go

CMD ["go", "run", "main.go"]
