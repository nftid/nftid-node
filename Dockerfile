FROM golang:alpine

RUN mkdir /node

ADD . /node

WORKDIR /node
RUN go mod download
RUN go build -o main .

CMD ["/node/main"]
