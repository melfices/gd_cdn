FROM golang:1.12
COPY . /go/src
EXPOSE 3000
WORKDIR /go/src
ENTRYPOINT [ "go", "run",  "main.go"]