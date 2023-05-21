FROM golang:1.16-alpine AS builder

RUN apk add --no-cache git

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

# Change the directory to cmd/main before building the image

WORKDIR /app/cmd/main/

RUN go build -o go-bookstore .

# Stage 2: Create the final image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/cmd/main/go-bookstore /app/cmd/main/go-bookstore

EXPOSE 9010

CMD ["/app/cmd/main/go-bookstore"]