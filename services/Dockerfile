FROM golang:1.8

ADD golang_echo_server.go /go/src/github.com/lavenderx/service_registry_discovery/

WORKDIR /go/src/github.com/lavenderx/service_registry_discovery

RUN go get github.com/labstack/echo/...
RUN go install

EXPOSE 1323

CMD ["go", "run", "golang_echo_server.go"]
