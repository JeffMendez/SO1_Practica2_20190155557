FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 6505

CMD ["go", "run", "main.go"]