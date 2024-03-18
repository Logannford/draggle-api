from golang:1.22.1

WORKDIR /

COPY . . 

RUN go build -o main *.go

CMD ["/main"]