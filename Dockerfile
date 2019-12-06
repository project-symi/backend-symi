FROM golang:1.13

RUN mkdir -p /src
ADD . /src
WORKDIR /src

RUN go build -o application application.go

CMD ["/src/application"]