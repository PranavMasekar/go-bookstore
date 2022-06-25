FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

# Change the directory to cmd/main before building the image

WORKDIR /app/cmd/main/

RUN go build -o go-bookstore .

EXPOSE 9010

CMD ["/app/cmd/main/go-bookstore"]