FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -u google.golang.org/grpc

RUN go build -o main .

EXPOSE 5505

CMD ["/app/main"]