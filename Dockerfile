FROM golang:1.22.1

WORKDIR /

COPY . . 

RUN go build -o main main.go

CMD ["/main"]